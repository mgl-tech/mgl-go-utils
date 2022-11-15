package sitemap

const (
	ChangeFreqAlways  = "always"
	ChangeFreqHourly  = "hourly"
	ChangeFreqDaily   = "daily"
	ChangeFreqWeekly  = "weekly"
	ChangeFreqMonthly = "monthly"
	ChangeFreqYearly  = "yearly"
	ChangeFreqNever   = "never"
)

/*type Url interface {
	SitemapLoc() string
	SitemapChangeFreq() string
	SitemapLastMod() string
	SitemapPriority() string
}*/

type Url struct {
	Loc        string
	ChangeFreq string
	LastMod    string
	Priority   string
	Language   string
}

/*func (u url) SitemapLoc() string {
	return u.Loc
}

func (u url) SitemapChangeFreq() string {
	return u.ChangeFreq
}

func (u url) SitemapLastMod() string {
	return u.LastMod
}

func (u url) SitemapPriority() string {
	return u.Priority
}
*/
