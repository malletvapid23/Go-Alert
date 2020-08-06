package smoketest

import (
	"encoding/json"
	"fmt"
	"github.com/target/goalert/smoketest/harness"
	"testing"
)

// TestNotices tests notices are properly returned given the appropriate scenarios.
func TestNotices(t *testing.T) {
	t.Parallel()

	const sql = `
		insert into escalation_policies (id, name) 
		values
			({{uuid "eid"}}, 'esc policy');
	`

	h := harness.NewHarness(t, sql, "contact-method-metadata")
	defer h.Close()

	doQL := func(query string, res interface{}) {
		g := h.GraphQLQuery2(query)
		for _, err := range g.Errors {
			t.Error("GraphQL Error:", err.Message)
		}
		if len(g.Errors) > 0 {
			t.Fatal("errors returned from GraphQL")
		}
		t.Log("Response:", string(g.Data))
		if res == nil {
			return
		}
		err := json.Unmarshal(g.Data, &res)
		if err != nil {
			t.Fatal("failed to parse response:", err)
		}
	}

	epID := h.UUID("eid")
	var notices struct {
		EscalationPolicy struct {
			Notices []struct {
				Type    string
				Message string
				Details string
			}
		}
	}

	// Verifying notice exists
	doQL(fmt.Sprintf(`
		query {
  			escalationPolicy(id: "%s") {
    			notices {
					type
					message
					details
    			}
  			}
		}
	`, epID), &notices)

	if len(notices.EscalationPolicy.Notices) < 1 {
		t.Fatalf("ERROR: retrieved length of notices = %d; want %d", len(notices.EscalationPolicy.Notices), 1)
	}
}
