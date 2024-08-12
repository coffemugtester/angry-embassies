package embassy_factory

func NewEmbassy(home, host, city string, consulate bool) *Embassy {
	embassy := &Embassy{
		HomeCountry: home,
		HostCountry: host,
		City:        city,
		Consulate:   consulate,
	}
	return embassy
}
