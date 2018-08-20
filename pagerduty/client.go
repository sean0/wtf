package pagerduty

import (
	"fmt"
	"github.com/PagerDuty/go-pagerduty"
	"os"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

type PolicyLevel struct {
	Users []string
}

type RestructuredEscalationPolicy struct {
	Levels map[string]PolicyLevel
}

func Fetch() (string, error) {
	var apikey = wtf.Config.UString("wtf.mods.pagerduty.apiKey", os.Getenv("WTF_PD_API_KEY"))
	if onCallsResponse, err := pagerDutyOnCallRequest(apikey); err != nil {
		return "", err
	} else {
		policiesMap := RestructureEscalationPolicies(onCallsResponse.OnCalls)
		return FormatOutput(policiesMap), nil
	}
}

func pagerDutyOnCallRequest(apiKey string) (*pagerduty.ListOnCallsResponse, error) {
	var opts pagerduty.ListOnCallOptions
	opts.UserIDs = wtf.ToStrs(wtf.Config.UList("wtf.mods.pagerduty.userIDs"))
	opts.EscalationPolicyIDs = wtf.ToStrs(wtf.Config.UList("wtf.mods.pagerduty.escalationPolicyIDs"))
	client := pagerduty.NewClient(apiKey)
	if onCallList, err := client.ListOnCalls(opts); err != nil {
		return nil, err
	} else {
		return onCallList, nil
	}
}

func RestructureEscalationPolicies(onCallsList []pagerduty.OnCall) map[string]RestructuredEscalationPolicy {
	policiesMap := make(map[string]RestructuredEscalationPolicy)
	var rep *RestructuredEscalationPolicy
	for _, onCall := range onCallsList {
		level := fmt.Sprint(onCall.EscalationLevel)
		var escPolName = onCall.EscalationPolicy.Summary
		var userName = onCall.User.Summary

		if escPol, ok := policiesMap[escPolName]; ok {
			if polLevel, ok := escPol.Levels[level]; ok {
				polLevel.Users = append(polLevel.Users, userName)
				policiesMap[escPolName].Levels[level] = polLevel
			} else {
				polLevel.Users = []string{userName}
				policiesMap[escPolName].Levels[level] = polLevel
			}
		} else {
			rep = &RestructuredEscalationPolicy{Levels: make(map[string]PolicyLevel)}
			if polLevel, ok := rep.Levels[level]; ok {
				polLevel.Users = append(polLevel.Users, userName)
				rep.Levels[level] = polLevel
			} else {
				rep.Levels[level] = PolicyLevel{Users: []string{userName}}
			}
			policiesMap[escPolName] = *rep
		}
	}
	return policiesMap
}

func FormatOutput(repMap map[string]RestructuredEscalationPolicy) string {
	var formattedOutput string
	for policyName, rep := range repMap {
		formattedOutput += fmt.Sprintf("[blue]Policy: %s[white]\n", policyName)
		for levelNumber, level := range rep.Levels {
			formattedOutput += fmt.Sprintf("[green]Level: %s - %s[white]\n", levelNumber, strings.Join(level.Users, ", "))
		}
		formattedOutput += fmt.Sprintf("\n")
	}
	return formattedOutput
}
