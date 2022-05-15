package xml_models

type Scope struct {
	Body struct {
		Text              string `xml:",chardata"`
		GetScopesResponse struct {
			Text   string `xml:",chardata"`
			Scopes []struct {
				Text      string `xml:",chardata"`
				ScopeDef  string `xml:"ScopeDef"`
				ScopeItem string `xml:"ScopeItem"`
			} `xml:"Scopes"`
		} `xml:"GetScopesResponse"`
	} `xml:"Body"`
}
