package prompt

import (
	"fmt"
	"testing"
)

// TestGenPrompt tests the GenPrompt function
func TestGenPrompt(t *testing.T) {
	prompt, err := GenPrompt(PromptData{
		DatabaseName: "MongoDB",
		Language:     "zh-CN",
		SlowQueryLog: "{op: 'command',ns: 'sample_mflix.comments',command: {aggregate: 'comments',pipeline: [{ '$match': { text: { '$regex': 'deserunt' } } },{ '$group': { _id: 1, n: { '$sum': 1 } } }],cursor: {},lsid: { id: UUI('20c0eb61-1798-49e7-9f42-d5f4706b161e') },'$db': 'sample_mflix'},keysExamined: 0,docsExamined: 50304,cursorExhausted: true,numYield: 50,nreturned: 1,queryHash: '8D6A792F',planCacheKey: '820BA36E',queryFramework: 'sbe',locks: {FeatureCompatibilityVersion: { acquireCount: { r: Long('52') } },Global: { acquireCount: { r: Long('52') } }},flowControl: {},storage: {},responseLength: 134,protocol: 'op_msg',cpuNanos: 43339083,millis: 43,planSummary: 'COLLSCAN',planningTimeMicros: 677,ts: ISODate('2024-07-14T08:10:17.427Z'),client: '192.168.65.1',appName: 'mongosh 2.2.11',allUsers: [],user: ''}",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("prompt:", prompt)
}

func Test_getSystemLanguage(t *testing.T) {
	fmt.Println(getSystemLanguage())
}
