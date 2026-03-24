# Data Sources for Synthetic Consumer Personas: Netherlands & Belgium

Deep research reference. Every source listed with: exact URL/API endpoint, auth requirements, data format, geographic granularity, temporal coverage, update frequency, limitations, licensing/terms.

---

## NETHERLANDS

### 1. Census / Demographics

#### CBS StatLine (Primary Source)
- **URL:** https://opendata.cbs.nl/statline/
- **API (OData v3):** `https://opendata.cbs.nl/ODataApi/odata/{TableID}` (metadata), `https://opendata.cbs.nl/ODataFeed/odata/{TableID}` (bulk data)
- **API (OData v4):** `https://datasets.cbs.nl/odata/v1/CBS/{TableID}/Observations`
- **Catalogue endpoint:** `https://opendata.cbs.nl/ODataCatalog/Tables`
- **Auth:** None required. Free, no registration.
- **Format:** JSON (OData), CSV download via StatLine portal, Excel download
- **Licence:** CC BY 4.0 (open data)
- **Update frequency:** Continuous (register-based census since 1971; no traditional decennial census)
- **Granularity:** Provincie (12), COROP (40), Gemeente (342), Wijk (~3,000), Buurt (~13,000)
- **Python client:** `pip install cbsodata` (v3) or `cbsodata4` (v4). R client: `cbsodataR`.
- **Gemeente codes:** GM prefix, e.g. GM0363 = Amsterdam

**Key table codes:**

| Table ID | Description | Granularity | Update |
|----------|-------------|-------------|--------|
| `03759ned` | Population by age, sex, migration background, gemeente | Gemeente | Annual (Jan 1) |
| `84910NED` | Population by migration background, generation, age, region (2010-2022) | Gemeente | Discontinued 2023 (replaced by herkomst tables) |
| `37296ned` | Population key figures | National/Provincie | Monthly |
| `83765NED` | Kerncijfers wijken en buurten (key figures by neighbourhood) | Wijk, Buurt | Annual |
| `85984NED` | Kerncijfers wijken en buurten 2024 | Wijk, Buurt | Annual |
| `86165NED` | Kerncijfers wijken en buurten 2025 | Wijk, Buurt | Annual |
| `82904NED` | Religieuze betrokkenheid (religious affiliation by personal characteristics) | National | Annual |
| `83288NED` | Religieuze betrokkenheid; kerkelijke gezindte; regio (religion by region) | Region | Periodic |
| `86052NED` | Education level (15-75) by wijk and buurt (2023) | Wijk, Buurt | Annual |
| `85525NED` | Education level by highest attainment and region | Provincie, COROP | Annual |
| `81955ENG` | Dwellings and non-residential stock (2012-2025) | Gemeente | Annual |
| `85035NED` | Woningvoorraad; woningtype, regio (housing stock by type) | Gemeente | Annual |
| `82900NED` | Housing stock; ownership, landlord type, occupancy, region | Gemeente | Annual |
| `85663NED` | CAO wages, contractual wage costs, working hours (2020=100 index) | Sector | Quarterly |
| `82838NED` | CAO wages (2010=100 historical, 1972-2023) | Sector | Annual |
| `84103ENG` | Income distribution of households (National Accounts) | National | Annual |
| `83906ENG` | Existing own homes; purchase prices, price indices (1995-2023) | National/Regional | Monthly |
| `83648NED` | Registered crime by type of crime and region | Politie-eenheid, Gemeente | Annual |
| `85146NED` | Veiligheidsmonitor kerncijfers by region | Politie-eenheid, 70k+ gemeente | Biennial (odd years) |
| `85844NED` | Veiligheidsmonitor kerncijfers by personal characteristics | National | Biennial (odd years) |
| `83678NED` | Budgetonderzoek: consumer spending by household type | National | ~5 yearly |

**Migration/herkomst note:** CBS changed terminology in 2022. "Migratieachtergrond" (Western/non-Western) replaced by "herkomst" (country of origin, born in NL yes/no). Tables 84910NED and 70751ned discontinued. New tables use country of birth of person and parents. Check catalogue for current herkomst tables.

#### CBS Kerncijfers Wijken en Buurten (Neighbourhood Key Figures)
- **URL:** https://www.cbs.nl/nl-nl/reeksen/publicatie/kerncijfers-wijken-en-buurten
- **Download:** Excel files per year (not possible to download all via StatLine simultaneously)
- **PDOK geodata:** https://www.pdok.nl/introductie/-/article/cbs-wijken-en-buurten (Wijk/Buurt boundaries as GeoJSON/WMS/WFS)
- **Kadaster Triple Store:** https://data.labs.kadaster.nl/cbs/cbs-wijken-buurten/
- **Content:** ~100 demographic, economic, housing, safety variables per wijk/buurt
- **Granularity:** Buurt (finest available, ~13,000 units)
- **Limitations:** Not all variables available at buurt level (suppression for small populations)

---

### 2. Consumer Spending

#### CBS Budgetonderzoek (Budget Survey)
- **Table:** 83678NED (via StatLine OData)
- **Methodology:** ~2,500 households record all expenses over 4 weeks. Annual since 1978 (gaps in 2001-2002).
- **Categories:** ~500 COICOP-aligned expenditure categories
- **Granularity:** National only (by household type, income decile, age group, not by region)
- **Format:** OData / CSV via StatLine
- **Update:** ~Every 5 years for full survey; annual micro sample
- **Microdata:** Available via CBS Remote Access Environment (RA) for registered researchers. Apply at https://www.cbs.nl/en-gb/our-services/customised-services-microdata/microdata-conducting-your-own-research
- **Limitations:** National-level aggregates only in open data. Regional spending data not published.

#### NIBUD Reference Budgets (Referentiebegrotingen)
- **URL:** https://www.nibud.nl/samenwerken/cijfers-en-rekentools/referentiebegrotingen/
- **Content:** Minimum and average household budgets by household type (single, couple, families with 1-4 children, elderly). Categories: housing, food, clothing, insurance, transport, telecom, leisure.
- **Access:** Published reports (free). Detailed budget tool licensed to organisations (paid).
- **Format:** PDF reports, some Excel tables
- **Granularity:** National (by household type, not by region)
- **Update:** Annual
- **Limitations:** Reference budgets, not actual spending data. Proprietary detailed data.

#### Eurostat HBS (Household Budget Survey)
- **URL:** https://ec.europa.eu/eurostat/web/microdata/household-budget-survey
- **Coverage:** NL data included. COICOP 3-digit classification.
- **Rounds:** 1988, 1994, 1999, 2005, 2010, 2015, 2020 (next: 2026)
- **Microdata:** Apply via Eurostat for scientific use (regulation EC 557/2013)
- **Aggregates:** Free at https://ec.europa.eu/eurostat/databrowser (table `hbs_exp_t121` etc.)
- **Granularity:** National (some NUTS1 breakdowns in aggregates)
- **Limitations:** 5-year cycle. Microdata requires formal application.

---

### 3. Media Consumption

#### NMO (Nationaal Media Onderzoek)
- **URL:** https://www.nationaalmediaonderzoek.nl/
- **Data portal:** https://nmodata.nl/
- **Composition:** Joint initiative of SKO (TV), NLO (radio), NOM (print), VINEX (internet). From 2025: includes out-of-home.
- **Coverage:** TV/video viewing (minutes/day), radio/audio, print readership, digital reach, cross-media
- **Access:** Summary data and press releases free. Detailed audience data: subscriber/member access only. Planning tools (proprietary) via NMO Data partners.
- **Format:** PDF reports, dashboards (subscriber-only)
- **Granularity:** National (by age, gender, SES)
- **Update:** Continuous measurement; quarterly/annual reports
- **Limitations:** Detailed data is proprietary. Only summary findings publicly available.

#### NOM (Nationaal Onderzoek Multimedia) - Print Readership
- **URL:** https://www.nommedia.nl/en
- **Content:** Print & Merken Monitor (NPMM) -- quarterly printreach figures for newspapers and magazines
- **Access:** Member/subscriber access for detailed data. Summary findings published as press releases.
- **Limitations:** Proprietary.

#### Newcom Research - Social Media Monitor
- **URL:** Reports via https://blossomyourcontent.eu/research-social-media-in-the-netherlands-2025/
- **Methodology:** Annual survey of ~6,400 Dutch people aged 15+
- **Content:** Platform-by-platform user counts, daily active users, growth/decline trends
- **Key 2025 findings:** WhatsApp 13.5M, YouTube 10.2M, Facebook 9.9M, Instagram 8.2M, TikTok growing, X declining (~500k lost)
- **Access:** Summary free. Full report: paid (via BlossomYourContent or Newcom)
- **Limitations:** Paid for detailed data. Annual snapshot, not continuous.

#### DataReportal Digital 2025: Netherlands
- **URL:** https://datareportal.com/reports/digital-2025-netherlands
- **Content:** Internet penetration, social media users by platform, device usage, e-commerce
- **Key stats:** Social media = 80.9% of total population; 13.1M users 18+
- **Access:** Free online viewing. SlideShare embeds.
- **Limitations:** Not downloadable as structured data.

#### Reuters Institute Digital News Report
- **URL (NL):** https://reutersinstitute.politics.ox.ac.uk/digital-news-report/2025/country-and-market-data (select Netherlands)
- **Full PDF:** https://reutersinstitute.politics.ox.ac.uk/sites/default/files/2025-06/Digital_News-Report_2025.pdf
- **NL partner:** Commissariaat voor de Media
- **Content:** News brand reach, trust in news, platform usage for news, willingness to pay
- **Access:** Free PDF download, free interactive charts
- **Format:** PDF, interactive web tool
- **Granularity:** National (by demographics)
- **Update:** Annual (June)

#### Regional Newspapers by Province

| Province(s) | Newspaper | Owner | Notes |
|-------------|-----------|-------|-------|
| Friesland | Leeuwarder Courant | NDC Mediagroep | Founded 1752, oldest Dutch daily |
| Friesland | Friesch Dagblad | NDC | Christian-oriented |
| Groningen, Drenthe | Dagblad van het Noorden | NDC | Merged 2002 |
| Overijssel | De Stentor | Mediahuis | Zwolle, Deventer, Apeldoorn editions |
| Gelderland | De Gelderlander | DPG Media | Nijmegen, Arnhem |
| Noord-Brabant | Brabants Dagblad | DPG Media | 's-Hertogenbosch, Tilburg |
| West-Brabant | BN/De Stem | DPG Media | Breda, Roosendaal |
| Limburg | De Limburger | Mediahuis | Only Limburg daily |
| Zeeland | Provinciale Zeeuwse Courant (PZC) | DPG Media | |
| Noord-Holland | Noordhollands Dagblad | HDC Media / Mediahuis | Alkmaar, Hoorn |
| Zuid-Holland | AD (regional editions) | DPG Media | Rotterdam, Den Haag, Leiden editions |
| Utrecht | AD Utrechts Nieuwsblad | DPG Media | Regional edition of AD |
| Flevoland | De Stentor (Flevoland ed.) | Mediahuis | |

**RSS notes:** Most Dutch regional newspapers (DPG Media titles: AD, Volkskrant, Trouw, Parool, Brabants Dagblad, PZC, Gelderlander, BN/De Stem, Tubantia) had RSS feeds but many have been discontinued or moved behind paywalls. Check individual sites.

#### National News RSS Feeds

| Source | RSS URL | Language |
|--------|---------|----------|
| NOS Nieuws Algemeen | `http://feeds.nos.nl/nosnieuwsalgemeen` | NL |
| NOS Binnenland | `http://feeds.nos.nl/nosnieuwsbinnenland` | NL |
| NOS Buitenland | `http://feeds.nos.nl/nosnieuwsbuitenland` | NL |
| NOS Economie | `http://feeds.nos.nl/nosnieuwseconomie` | NL |
| NOS Sport | `http://feeds.nos.nl/nossportalgemeen` | NL |
| NOS Tech | `http://feeds.nos.nl/nosnieuwstech` | NL |
| NOS Cultuur & Media | `http://feeds.nos.nl/nosnieuwscultuurenmedia` | NL |
| NOS Politiek | `http://feeds.nos.nl/nosnieuwspolitiek` | NL |
| De Telegraaf | `https://www.telegraaf.nl/rss` | NL |
| NRC | `https://www.nrc.nl/rss/` | NL |
| De Volkskrant | `https://www.volkskrant.nl/voorpagina/rss` | NL |
| AD (Algemeen Dagblad) | `https://www.ad.nl/home/rss.xml` | NL |
| Trouw | `https://www.trouw.nl/voorpagina/rss.xml` | NL |
| Het Parool | `https://www.parool.nl/voorpagina/rss.xml` | NL |
| DutchNews.nl | `https://www.dutchnews.nl/feed/` | EN |
| NL Times | `https://nltimes.nl/feed` | EN |
| Government.nl | `https://www.government.nl/rss` | EN |

---

### 4. Social Attitudes / Institutional Trust

#### LISS Panel (Longitudinal Internet Studies for the Social Sciences)
- **URL:** https://www.lissdata.nl/
- **Operator:** CentERdata (Tilburg University), part of ODISSEI since 2019
- **Panel:** ~5,000 households (~7,500 individuals 16+), true probability sample from population register
- **Core study modules:** Repeated annually: health, politics, religion, social integration, economic, personality, work, education, family, values, housing
- **Access:** Free for academic research. Sign user statement at https://www.lissdata.nl. No commercial use.
- **Format:** SPSS (.sav), Stata (.dta)
- **Granularity:** Individual-level microdata; regional identifiers limited for privacy
- **Temporal:** Running since 2007, annual waves
- **Variables:** Hundreds per wave. Core: demographics, income, education, political preference, religiosity, social trust, media use, health
- **Limitations:** No commercial use. No geographic identifiers finer than province (privacy).

#### SCP (Sociaal en Cultureel Planbureau)
- **URL:** https://www.scp.nl/ (Dutch), https://english.scp.nl/ (English)
- **Content:** Publications on quality of life, social cohesion, trust, cultural participation, civic engagement
- **Data:** Reports (PDF, free download). Raw data not routinely published as open data. Some datasets deposited at DANS (https://dans.knaw.nl/).
- **Key reports:** "Sociale Staat van Nederland" (biennial), "Burgerperspectieven" (quarterly, trust/sentiment)
- **Limitations:** Mostly report-level findings, not raw data. Academic access via DANS for some datasets.

#### Eurobarometer
- **URL:** https://europa.eu/eurobarometer/surveys/browse/all
- **Microdata:** Free via GESIS (https://www.gesis.org/en/eurobarometer-data-service). SPSS/Stata format.
- **NL coverage:** Standard Eurobarometer twice yearly. Trust in institutions, EU attitudes, policy attitudes.
- **Granularity:** National (NL). No sub-national breakdown.
- **Temporal:** Since 1973. ~100 waves.
- **Auth:** Register at GESIS (free, academic/research purpose statement)

#### OECD Trust Survey 2024
- **NL country note:** https://www.oecd.org/en/publications/oecd-survey-on-drivers-of-trust-in-public-institutions-2024-results-country-notes_a8004759-en/netherlands_d71d7263-en.html
- **Microdata:** Apply at https://www.oecd.org/en/data/datasets/oecd-trust-survey-data.html (Terms of Use form required)
- **Content:** Trust in parliament, government, courts, police, media; satisfaction with public services
- **Limitations:** Application-based access. ~60,000 respondents across 30 OECD countries (NL sample ~2,000).

#### European Social Survey (ESS)
- **Data portal:** https://www.europeansocialsurvey.org/data-portal
- **NL participation:** Rounds 1-11 (2002-2023/24)
- **Content:** Social trust, political trust, immigration attitudes, subjective well-being, discrimination, media use
- **Access:** Free, register at portal. SPSS/Stata/CSV.
- **Granularity:** Individual-level; NUTS1 region available in some rounds
- **Limitations:** ~1,500-2,000 NL respondents per round.

#### ISSP (International Social Survey Programme)
- **Data:** Free via GESIS (https://www.gesis.org/en/issp/data-and-documentation)
- **NL participation:** Yes (multiple modules)
- **Modules:** Religion, Role of Government, Social Inequality, National Identity, Environment, Work Orientations, Health, Social Networks, Citizenship, Family
- **Access:** Free download. SPSS/Stata.
- **Granularity:** National only
- **Limitations:** 15-minute supplement to national surveys. Rotating module every ~10 years.

---

### 5. Political Landscape

#### Kiesraad (Electoral Council) - Official Election Results
- **URL:** https://www.kiesraad.nl/ (NL), https://english.kiesraad.nl/ (EN)
- **Results portal:** https://www.verkiezingsuitslagen.nl/
- **Data download:** Election Markup Language (EML) files at https://data.overheid.nl/community/organization/kiesraad
- **GitHub:** https://github.com/kiesraad (maps, tools)
- **Third party scraper:** https://github.com/DIRKMJK/kiesraad (Python)
- **Format:** EML (XML-based), CSV (per gemeente from results pages), GeoJSON (maps)
- **Auth:** None
- **Granularity:** Gemeente (municipality), per polling station for recent elections
- **Temporal:** Database covers 700+ elections back to 1848
- **Recent elections:**
  - Tweede Kamer: 22 November 2023
  - Provinciale Staten: 15 March 2023
  - Gemeenteraad: 16 March 2022
  - Eerste Kamer: 30 May 2023 (indirect)
  - Europees Parlement: 6 June 2024

#### Peilingwijzer (Polling Aggregator)
- **URL:** https://www.tomlouwerse.nl/project/polling-aggregation/
- **GitHub:** https://github.com/louwerse/peilingwijzer (replication materials)
- **Creator:** Prof. Tom Louwerse, Leiden University (published by NOS)
- **Methodology:** Bayesian smoothing of polls from I&O Research, Ipsos, Peil, LISS panel, De Stemming. Adjusts for house effects.
- **Access:** Published aggregates on NOS. Replication code on GitHub.
- **Format:** No structured download (chart/visual output). Replication materials in R.
- **Limitations:** Academic project, not a structured data feed.

#### Political Geography Notes
- **Bible Belt:** Conservative Reformed. Band from Zeeland through Goeree-Overflakkee, Alblasserwaard, Veluwe to Overijssel. SGP (Reformed Political Party) stronghold.
- **Randstad:** Urban ring (Amsterdam, Rotterdam, The Hague, Utrecht). D66, GroenLinks-PvdA strong in Amsterdam/Utrecht.
- **Limburg:** PVV historically strong. Catholic tradition.
- **Southern Catholic:** Noord-Brabant and Limburg. CDA historically; now fragmented.
- **Northern Protestant:** Groningen, Friesland, Drenthe. Mix of PvdA tradition and BBB rural support.
- **Rural provinces:** BBB breakthrough in 2023 Provinciale Staten.

---

### 6. Employment / Salary

#### CBS Employment and Earnings Statistics (SWL)
- **Coverage:** All employee jobs under Dutch wage tax legislation
- **Source data:** Wage declarations to tax authority (since 2006)
- **Key tables:**
  - Income by personal characteristics: search StatLine for "Inkomen van personen"
  - Income distribution by gemeente: search for "Inkomen" + region tables
  - `84583NED`: Wages by sector and occupation
- **Granularity:** National by sector/occupation; income by gemeente in some tables
- **Update:** Annual
- **Limitations:** Self-employed income separate. Occupation-by-region cross-tabs limited.

#### UWV (Uitvoeringsinstituut Werknemersverzekeringen) Labour Market Data
- **URL:** https://www.uwv.nl/nl/arbeidsmarktinformatie/datasets
- **Alt URL:** https://www.werk.nl/arbeidsmarktinformatie/datasets
- **Content:** Vacancies, unemployment (WW), labour market tension indicator, by region/profession/age/gender
- **Format:** Excel/CSV download
- **Auth:** None
- **Granularity:** Arbeidsmarktregio (35 labour market regions), UWV-district
- **Update:** Monthly (third Thursday)
- **Limitations:** Focuses on unemployment/vacancies rather than salary levels.

#### CAO Wage Data
- **CBS tables:** 85663NED (2020=100 index), 82838NED (historical 1972-2023)
- **Content:** Contractual wages, wage costs, working hours from 208 major CAOs covering ~95% of CAO employees
- **SZW register:** Ministry of Social Affairs publishes CAO texts at https://open.overheid.nl/ (search "CAO")
- **Limitations:** Index numbers, not absolute salary levels by occupation/region.

---

### 7. Housing Market

#### CBS Housing Tables
- **Key tables:** 85035NED (stock by type/region), 82900NED (ownership/landlord type), 81955ENG (dwellings 2012-2025)
- **Content:** Owner-occupied vs rental (woningcorporatie vs private), dwelling type, construction year, WOZ value bands
- **Granularity:** Gemeente
- **Update:** Annual (January 1 snapshot)

#### Kadaster / PDOK
- **URL:** https://www.kadaster.nl/zakelijk/datasets/open-datasets
- **PDOK:** https://www.pdok.nl/ (WMS/WFS/API for cadastral boundaries, BAG address register)
- **BAG (Basisregistratie Adressen en Gebouwen):** Free, full address + building register via PDOK API
- **Property transaction prices:** PAID dataset from Kadaster. Not in open data.
- **WOZ values:** Public information per property, but bulk access not open. CBS publishes aggregate WOZ by gemeente.
- **Limitations:** Individual transaction prices and WOZ values require paid Kadaster subscription or CBS microdata access.

#### NVM (Nederlandse Vereniging van Makelaars)
- **URL:** https://www.nvm.nl/
- **Content:** Quarterly price indices, transaction volumes, asking vs selling prices
- **Access:** Summary press releases free. Detailed data: members only (proprietary).
- **Limitations:** Proprietary. NVM covers ~75% of transactions (not all).

---

### 8. Crime / Safety

#### CBS Veiligheidsmonitor (Safety Monitor)
- **URL:** https://www.cbs.nl/nl-nl/maatschappij/veiligheid-en-recht/veiligheidsmonitor
- **Latest:** Veiligheidsmonitor 2025 (~200,000 respondents aged 15+)
- **Tables:** 85146NED (by region), 85844NED (by person characteristics)
- **Content:** Neighbourhood safety perception, nuisance, victimisation (traditional + online crime), police contact, crime prevention
- **Granularity:** Province, 70k+ municipalities (G4, G40), police districts, regional units
- **Frequency:** Biennial (odd years) since 2017; annual before that
- **Limitations:** Regional detail limited to larger municipalities. Buurt/wijk-level not in open data.

#### Politie Open Data
- **URL:** https://data.politie.nl/
- **CBS crime tables:** 83648NED (registered crime by type and region)
- **Data.overheid:** "Geregistreerde misdrijven en aangiften; soort misdrijf, gemeente" for gemeente-level
- **Content:** Registered offences by type, reportings, per 1,000 residents
- **Format:** Excel download (per gemeente or national)
- **Granularity:** Gemeente, politie-eenheid (regional unit)
- **Update:** Monthly (police data), annual (CBS aggregates)

---

### 9. Religion

#### CBS Religion Tables
- **Table 82904NED:** Religious affiliation by personal characteristics (age, sex, education)
- **Table 83288NED:** Religious affiliation by region
- **Categories (post-2019):** No religion, Roman Catholic, Protestant (PKN merged), Islam, Other
- **Categories (pre-2019):** Distinguished Dutch Reformed (Hervormd), Reformed (Gereformeerd) separately
- **Update:** Annual. Figures for 2024 added April 2025; 2025 expected H1 2026.
- **SCP reports:** "Religie in Nederland" series (PDF, free)
- **Limitations:** Regional religion data is province-level, not gemeente.

#### Regional Religion Patterns
- **Bible Belt:** Zeeland (Schouwen, Goeree), South Holland (Alblasserwaard), Gelderland (Veluwe), Overijssel (Staphorst, Urk). SGP voter maps proxy for strict Reformed communities.
- **Catholic south:** Noord-Brabant, Limburg
- **Secular cities:** Amsterdam, Rotterdam, The Hague, Utrecht (60-70% no religion)
- **Muslim concentrations:** Rotterdam-Zuid, Amsterdam-West/Nieuw-West, The Hague Schilderswijk, Utrecht Kanaleneiland

---

### 10. Health

#### RIVM Gezondheidsmonitor (Health Monitors)
- **Open data:** https://www.rivm.nl/media/smap/opendata.html
- **StatLine portal:** https://statline.rivm.nl/
- **Content:** Health, well-being, lifestyle of Dutch population. Three monitors: youth (0-11), young adults (12-25), adults (19+) and elderly (65+).
- **Conducted by:** GGD'en (25 Municipal Health Services), CBS, RIVM jointly
- **Granularity:** GGD-regio (25), Gemeente, Wijk (estimates with confidence intervals)
- **Variables:** BMI, smoking, alcohol, physical activity, loneliness, psychological distress, chronic conditions, healthcare use
- **Format:** CSV/Excel download, OData via RIVM StatLine
- **Update:** Every 4 years (next adult monitor ~2024/2025)
- **Limitations:** Wijk-level estimates have wide confidence intervals for small populations.

#### CBS Gezondheidsenquete (Health Survey)
- **Table codes:** Search StatLine for "Gezondheid" tables
- **Content:** Self-reported health, medical contacts, lifestyle, preventive behaviour
- **Running:** Annually since 1981
- **Microdata:** Available via DANS (https://dans.knaw.nl/) as "Gezondheidsenquete 2014-2023" or CBS Remote Access
- **Limitations:** Open data is aggregate; individual-level requires researcher application.

#### Lifestyle Monitor
- **URL:** https://www.rivm.nl/en/lifestyle-monitor
- **Content:** Comprehensive lifestyle tracking (nutrition, physical activity, smoking, alcohol, drugs)
- **Partners:** RIVM, CBS, GGD'en, Voedingscentrum, Trimbos-instituut, Rutgers
- **Access:** Online dashboard, PDF reports
- **Granularity:** GGD-regio, some gemeente-level

---

### 11. Income by Region

#### CBS Income Tables
- **Key tables:**
  - Search StatLine for "Inkomen van huishoudens": household income by municipality
  - `85039ned`: Income distribution by gemeente (part of Kerncijfers wijken en buurten)
  - Income deciles/quintiles available at gemeente level
- **Content:** Standardised household income, personal income, income sources (employment, benefits, pensions)
- **Granularity:** Gemeente, Wijk (in Kerncijfers tables)
- **Heat maps:** https://www.cbs.nl/en-gb/about-us/innovation/project/heat-maps-with-income-level-of-men-and-women (100m grid)

---

---

## BELGIUM

### 1. Census / Demographics

#### StatBel (Statistics Belgium) Open Data
- **Portal:** https://statbel.fgov.be/en/open-data
- **Alt portal:** https://data.gov.be/en (Belgian federal open data portal)
- **Auth:** None required
- **Licence:** CC BY 4.0
- **Format:** XLSX, CSV, TXT (varies by dataset). Some datasets on data.gov.be have API access.
- **CRITICAL:** Data in bilingual column names (NL/FR). Some datasets only available in one language.
- **No formal REST API** for StatBel itself. Bulk download approach. The data.gov.be portal has some datasets with CKAN/DCAT API.

**Key datasets:**

| Dataset | URL Slug | Granularity | Format | Update |
|---------|----------|-------------|--------|--------|
| Population by statistical sector | `/open-data/population-statistical-sector-2024` | Statistical sector (~20,000) | XLSX/CSV | Annual |
| Population by place of residence, nationality, marital status, age, sex | data.gov.be "statbel-population-2023" | Gemeente/Commune (581) | CSV | Annual |
| Fiscal statistics on income | `/open-data/fiscal-statistics-income-statistical-sector` | Statistical sector | XLSX | Annual (2005-2023) |
| Real estate sales by municipality (2010-2017) | `/open-data/sales-real-estate-municipality-...` | Commune | XLSX | Periodic |
| Real estate sales (historical 1973-2017) | `/open-data/sales-real-estate-belgium-...` | Commune/Province | XLSX | Static |
| Statistical sectors 2020 (boundaries) | `/open-data/statistical-sectors-2020` | Statistical sector | SHP/GeoJSON | Periodic |
| Census 2011 tables | `/open-data` (search "Census") | Commune | XLSX | Static |
| Establishment units | `/open-data` (search "establishment") | Commune (2015-2023) | CSV | Annual |

**Administrative hierarchy:** Gewest/Region (3: Flanders, Wallonia, Brussels-Capital) > Provincie/Province (10) > Arrondissement (43) > Gemeente/Commune (581) > Statistical sector (~20,000)

**Population by nationality/origin:**
- StatBel publishes by nationality at birth and current nationality
- Belgium does not collect ethnic data. Proxy: country of birth + parents' country of birth
- Origin data: https://statbel.fgov.be/en/themes/population/structure-population/origin

#### BelgiumStatistics R Package
- **GitHub:** https://github.com/weRbelgium/BelgiumStatistics
- **Content:** R package wrapping StatBel open data + administrative boundary maps
- **Includes:** Population data, boundary polygons at all admin levels
- **Limitations:** Community-maintained, may lag behind StatBel releases.

---

### 2. Consumer Spending

#### StatBel Household Budget Survey (HBS)
- **URL:** https://statbel.fgov.be/en/themes/households/household-budget-survey-hbs
- **Data.gov.be:** https://data.gov.be/en/datasets/67b8a34d4c6981a19f4cad0e779df597e8981c95
- **Methodology:** Biennial survey of ~5,000-5,500 households
- **Classification:** COICOP (Classification of Individual Consumption by Purpose)
- **Content:** Average expenditure by category (housing, food, transport, clothing, leisure, etc.) by household type, income quintile, region
- **Format:** PDF report, Excel tables on StatBel site
- **Granularity:** Region (Flanders, Wallonia, Brussels), household type, income quintile. NOT commune-level.
- **Latest:** 2024 survey (results being published 2025)
- **Update:** Biennial
- **Limitations:** Regional-level only (3 regions). No sub-regional consumer spending data.

#### Eurostat HBS for Belgium
- Same as Netherlands section above. Belgium included. COICOP 3-digit. Microdata via application.

#### Flanders vs Wallonia vs Brussels Spending Differences
- StatBel HBS does break down by region. Key differences: Brussels housing costs highest, Wallonia transport costs higher (rural), Flanders higher discretionary spending.
- No open dataset provides finer than regional breakdown.

---

### 3. Media Consumption

#### CIM (Centre d'Information sur les Medias)
- **URL:** https://www.cim.be/en
- **Role:** Official Belgian audience measurement body since 1971. 150 members (advertisers, media, agencies).
- **Covers:** TV, radio, press, digital, out-of-home
- **Public data:** Daily results of press site audiences. Summary audience figures in press releases.
- **Detailed data:** PROPRIETARY. Available to CIM members via media planning tools (e.g. Gemius). Login at www.e.gemius.com/login.
- **Methodology:** ~8,000 people surveyed annually by IPSOS for press; TAM panel for TV; online measurement via Gemius
- **Limitations:** Granular audience data behind paywall. Only aggregate summaries public.

#### Reuters Institute Digital News Report - Belgium
- **URL:** https://reutersinstitute.politics.ox.ac.uk/digital-news-report/2025/belgium
- **BE research partner:** SMIT, Vrije Universiteit Brussel
- **Content:** News brand reach (split Flemish/Francophone), trust in news, platform usage, payment behaviour
- **Access:** Free PDF + interactive charts
- **Key insight:** Belgium has completely split media markets: Flemish and Francophone audiences consume entirely different news brands

#### DataReportal Digital 2025: Belgium
- **URL:** https://datareportal.com/digital-in-belgium
- **Content:** Internet users, social media users by platform, device usage
- **Access:** Free viewing online

#### Belgian Newspaper Landscape

**Flemish (Dutch-language):**

| Newspaper | Owner | Region/Focus |
|-----------|-------|-------------|
| Het Laatste Nieuws (HLN) | DPG Media | Popular, Flanders-wide, largest circulation |
| De Standaard | Mediahuis | Quality broadsheet, Flanders-wide |
| De Morgen | DPG Media | Progressive, Flanders-wide |
| Het Nieuwsblad | Mediahuis | Popular, Flanders-wide, strong local editions |
| Gazet van Antwerpen | Mediahuis | Antwerp province |
| Het Belang van Limburg | Mediahuis | Belgian Limburg province |

**Francophone (French-language):**

| Newspaper | Owner | Region/Focus |
|-----------|-------|-------------|
| Le Soir | Rossel | Quality broadsheet, Brussels/Wallonia |
| La Libre Belgique | IPM | Centre-right, Brussels/Wallonia |
| La Derniere Heure (DH/Les Sports) | IPM | Popular, Brussels/Wallonia |
| L'Avenir | Nethys/Publifer | Regional editions across Wallonia (Namur, Liege, Luxembourg, etc.) |
| L'Echo | Mediafin (Rossel/Roularta) | Financial daily |
| La Meuse | Rossel | Liege region |
| La Nouvelle Gazette | Rossel | Hainaut/Charleroi |
| La Province | Rossel | Mons/Hainaut |
| Nord Eclair | Rossel | Tournai/Hainaut |

#### Belgian News RSS Feeds

| Source | RSS URL | Language |
|--------|---------|----------|
| VRT NWS | `https://www.vrt.be/vrtnws/nl.rss.articles.xml` | NL |
| HLN (Het Laatste Nieuws) | `https://www.hln.be/home/rss.xml` | NL |
| De Morgen | `https://www.demorgen.be/in-het-nieuws/rss.xml` | NL |
| L'Echo | `https://www.lecho.be/rss/top_stories.xml` | FR |
| RTBF Info | `https://rss.rtbf.be/article/rss/rtbfinfo_homepage.xml` | FR |
| De Standaard | `https://www.standaard.be/rss/section/1f2838d4-99ea-49f0-9102-138784c7ea7c` (front page) | NL |
| Le Soir | `https://www.lesoir.be/rss/homepage.xml` | FR |
| La Libre | `https://www.lalibre.be/arc/outboundfeeds/rss/section/belgique/` | FR |
| Brussels.be | `https://www.brussels.be/rss-feeds` (multiple feeds) | NL/FR/EN |

**Note:** Belgian newspaper RSS feeds change frequently. Verify before implementing.

---

### 4. Social Attitudes / Institutional Trust

#### Eurobarometer - Belgium
- Same source as NL (GESIS). Belgium included in all Standard Eurobarometer waves.
- **Key insight:** Belgium usually shows lower trust in national institutions than NL; significant Flanders/Wallonia split.

#### OECD Trust Survey 2024 - Belgium
- **Country note:** https://www.oecd.org/en/publications/oecd-survey-on-drivers-of-trust-in-public-institutions-2024-results-country-notes_a8004759-en/belgium_fb2afb08-en.html
- Same application process as NL for microdata.

#### European Social Survey (ESS) - Belgium
- Same portal as NL. Belgium has participated in Rounds 1-11.
- **Key value:** Contains Flanders/Wallonia/Brussels NUTS2 identifiers in some rounds, enabling regional attitudinal analysis.

#### ISSP - Belgium
- Belgium participates in ISSP via GESIS. Same modules as NL section above.

#### PartiRep / RepResent (Belgian Election Studies)
- **PartiRep I (2007-2011):** Electoral survey around 2009 regional elections
  - Data: https://easy.dans.knaw.nl/ui/datasets/id/easy-dataset:262840
- **PartiRep II (2012-2017):** Electoral survey around 2014 elections
  - Data: https://easy.dans.knaw.nl/ui/datasets/id/easy-dataset:262841
- **RepResent Panel (2019-2021):** 4-wave voter panel around May 2019 elections
  - PIs: KU Leuven, UCLouvain, VUB
  - ~250 variables per wave. SPSS/Stata.
  - Access: via project PIs (apply for data access)
- **Belgian National Elections Study (BNES):** 1991-2019
  - Institute: ISPO, KU Leuven (https://soc.kuleuven.be/ceso/ispo)
  - Access: Apply via researchers

#### EU-SILC Belgium
- **Operator:** StatBel (Belgian implementation)
- **Survey:** ~6,000 households (~11,000 people) per year, 4-year panel
- **Content:** Income, poverty, social exclusion, living conditions, housing costs, child care, health
- **Regional data:** Breakdown by Flanders/Wallonia/Brussels available in published indicators
- **Microdata:** Apply via Eurostat (regulation 557/2013)
- **Published indicators:** At https://statbel.fgov.be/en/themes/households/poverty-and-living-conditions

---

### 5. Political Landscape

#### SPF Interieur / FOD Binnenlandse Zaken - Official Election Results
- **Results portal:** https://electionresults.belgium.be/ (verkiezingsresultaten.belgium.be)
- **JSON API:** https://api.electionresults.belgium.be/swagger/index.html
  - Endpoints for results by commune, canton, province, region
  - Returns JSON
- **Download:** From results portal, click download icon (top right), choose CSV, XML, or JSON
- **Auth:** None
- **Alternative API:** Open Knowledge Belgium: https://github.com/jbelien/elections-be-api (historical data)
- **Granularity:** Commune (581), Canton, Province, Region, Federal
- **Elections covered:**
  - Federal: 2024 (9 June), 2019, 2014, 2010, ...
  - Regional: 2024 (9 June), 2019, 2014, ...
  - European: 2024 (9 June), 2019, ...
  - Local/communal: 2024 (13 October), 2018, ... (note: local elections transferred to communities/regions post-2001)
- **Compulsory voting:** Turnout typically 87-92%

#### Split Party System
- **Flemish parties:** N-VA, Vlaams Belang, Vooruit, Open Vld, CD&V, Groen
- **Francophone parties:** PS, MR, Les Engages (ex-cdH), Ecolo, DeFI
- **Bilingual (only one):** PTB/PVDA (Marxist)
- **No national parties:** Voters can only vote for parties in their linguistic constituency
- **Coalition complexity:** Federal government requires majority in both language groups

#### Europe Elects - Belgium Polling
- **URL:** https://europeelects.eu/belgium/
- **Content:** Aggregated Belgian polling data, seat projections
- **Access:** Free

---

### 6. Employment / Salary

#### StatBel Wage Statistics
- **Overview:** https://statbel.fgov.be/en/themes/work-training/wages-and-labourcost/overview-belgian-wages-and-salaries
- **Gross wages:** https://statbel.fgov.be/en/themes/indicators/labour/gross-wages-and-salaries
- **Content:** Average monthly wage by sector, occupation (ISCO-08), region, sex
- **Key data:** Brussels-Capital average = EUR 4,748/month (16% above national average). Chief executives: EUR 11,772 (189% above average).
- **Open data (fiscal):** Fiscal statistics on income by statistical sector (2005-2023): https://statbel.fgov.be/en/open-data/fiscal-statistics-income-statistical-sector
- **Format:** XLSX
- **Granularity:** Statistical sector (~20,000) for fiscal income; Province/Region for wages
- **Licence:** CC BY 4.0

#### VDAB (Flemish Employment Service)
- **Statistics:** https://www.vdab.be/trends/arvastat
- **Content:** Unemployment by municipality, vacancies by sector/region, labour market tension
- **Granularity:** Gemeente (Flemish municipalities)
- **Format:** Online dashboard, some Excel downloads
- **Access:** Free
- **Limitations:** Flanders only

#### FOREM (Walloon Employment Service)
- **Statistics:** https://www.leforem.be/chiffres-et-analyses-du-marche-de-l-emploi.html
- **Content:** Unemployment, vacancies, sector analysis for Wallonia
- **Granularity:** Commune (Walloon)
- **Format:** Online reports, PDF
- **Access:** Free
- **Limitations:** Wallonia only

#### Actiris (Brussels Employment Service)
- **Statistics:** https://www.actiris.brussels/nl/Burgers/cijfers
- **Content:** Brussels-Capital Region employment data
- **Granularity:** Brussels-Capital Region
- **Format:** Online reports
- **Access:** Free
- **Limitations:** Brussels only

#### Synerjob (Cross-regional Coordination)
- **URL:** https://www.synerjob.be/en/statistieken.html
- **Content:** Links to statistics from all three regional employment agencies

---

### 7. Housing Market

#### StatBel Real Estate
- **URL:** https://statbel.fgov.be/en/themes/housing/real-estate
- **Open data downloads:**
  - Sales by municipality (2010-2017): XLSX at `/open-data/sales-real-estate-municipality-according-nature-property-land-register-2010-2017`
  - Historical series (1973-2017): similar URL pattern
  - Transaction count, total price, median price, percentiles (P10, Q25, Q50, Q75, P90)
- **Source:** Deeds of sale from FPS Finances (GAPD / land registry)
- **Format:** XLSX
- **Granularity:** Commune/Municipality
- **Licence:** CC BY 4.0
- **Limitations:** Most recent open data ends 2017 for detailed municipal data. More recent data in reports/indicators but may not be downloadable at commune level.

#### Belgian Notaries (Fednot)
- **URL:** https://www.notaire.be/ (FR) / https://www.notaris.be/ (NL)
- **Content:** Quarterly barometer of real estate prices
- **Granularity:** Province, region
- **Access:** Press releases (free). Detailed data proprietary.

---

### 8. Crime / Safety

#### Belgian Federal Police Statistics
- **Portal:** https://www.police.be/statistiques/fr (FR) / https://www.police.be/statistiques/nl (NL)
- **Content:** Registered offences by type, traffic accidents, police morphology data
- **Granularity:** Commune, Police zone (188 zones), Judicial district, Province, Region, Federal
- **Access:** Interactive dashboards with tables, graphs, maps. Downloadable.
- **Format:** Excel/CSV from dashboard tool
- **Auth:** None
- **Update:** Annual (main report); some data quarterly
- **Limitations:** Not all crime types available at commune level (suppression for small numbers).

---

### 9. Religion

#### Officially Recognised Denominations
Belgium recognises 6 religions + organised laicite:
1. Roman Catholicism
2. Protestantism (United Protestant Church)
3. Anglicanism
4. Orthodox Christianity
5. Islam
6. Judaism
7. Non-religious philosophical organisations (laicite organisee)
Buddhism recognition pending under secular organisation framework.

#### Religious Data Sources
- **No census question on religion** since 1846. All data from surveys.
- **European Social Survey (ESS):** Religious affiliation and attendance, Belgium included. Regional breakdown (Flanders/Wallonia/Brussels) in some rounds.
- **Eurobarometer:** Periodic religion questions
- **Catholic Church in Belgium annual report:** Mass attendance figures (173,335 regular Sunday attendance in October 2024, slight increase from 2023)
- **Regional patterns:** Catholic practice declining faster in Wallonia than Flanders. Brussels heavily secularised but large Muslim community (~24% of Brussels population). Flanders: 50% Catholic identification (2022), declining. Weekly mass attendance ~7% in Flanders.
- **Limitations:** No official statistics. All figures are survey estimates or church self-reports.

---

### 10. Health

#### Sciensano - Health Interview Survey (HIS)
- **URL:** https://www.sciensano.be/en/projects/health-interview-survey
- **Interactive tool:** HISIA (online analysis tool)
- **Latest:** 7th HIS (2023-2024), 7,001 participants
- **Content:** Health status, lifestyle, health care use, mental health, chronic conditions, preventive behaviour
- **Granularity:** Region (Flanders, Wallonia, Brussels), Province, limited commune-level
- **Microdata:** Apply to Sciensano HIS team for anonymised microdata (scientific institutions)
- **Published data:** Infographics, reports, HISIA tool (free)
- **Format:** PDF reports, interactive online tool
- **Data.gov.be:** https://data.gov.be/en/datasets/79643855-6a56-4604-91f4-e92728afd54d
- **Update:** Every ~5 years (previous: 2018, 2013, 2008, 2004, 2001, 1997)

#### RIZIV/INAMI (National Institute for Health and Disability Insurance)
- **URL:** https://www.riziv.fgov.be/ (NL) / https://www.inami.fgov.be/ (FR)
- **Content:** Healthcare utilisation data, reimbursement statistics, pharmaceutical consumption
- **Granularity:** Province/region for aggregate stats; individual-level via linked datasets (research application)
- **Access:** Aggregate data in annual reports. Detailed data requires formal application.
- **Limitations:** Administrative data, not survey-based health status data.

---

### 11. Additional Cross-Cutting Sources

#### Eurostat (applies to both NL and BE)
- **Data browser:** https://ec.europa.eu/eurostat/databrowser
- **NUTS classification:** NL = NUTS0; provinces = NUTS2; COROP = NUTS3. BE = NUTS0; regions = NUTS1; provinces = NUTS2.
- **Key datasets:**
  - `nama_10_co3_p3`: COICOP household consumption
  - `ilc_*`: Income and living conditions (EU-SILC)
  - `lfsa_*`: Labour force survey
  - `demo_*`: Demographics
  - `hlth_*`: Health
- **Format:** TSV, CSV, JSON, SDMX via bulk download or API
- **Auth:** None for aggregates. Microdata requires application.
- **Granularity:** NUTS2 for most datasets; some NUTS3

#### GESIS Data Archive
- **URL:** https://www.gesis.org/en/eurobarometer-data-service
- **Content:** Eurobarometer, ISSP, ESS data. Free registration required.
- **Format:** SPSS, Stata
- **Licence:** Academic use (non-commercial for most datasets)

#### European Values Study (EVS)
- **URL:** https://europeanvaluesstudy.eu/
- **NL + BE coverage:** Multiple waves (1981, 1990, 1999, 2008, 2017)
- **Content:** Values, attitudes, beliefs across European societies
- **Access:** GESIS data archive
- **Granularity:** National (with region in some waves)

---

## SUMMARY: Actionable Data by Category

### Netherlands -- Finest Granularity Available

| Category | Source | Finest Granularity | Open/Free |
|----------|--------|--------------------|-----------|
| Population demographics | CBS StatLine | Buurt (~13,000) | Yes (OData API) |
| Migration/origin | CBS StatLine | Gemeente (342) | Yes |
| Education level | CBS StatLine | Wijk/Buurt | Yes |
| Income distribution | CBS StatLine / Kerncijfers | Wijk | Yes |
| Consumer spending | CBS Budgetonderzoek | National only | Open (aggregates) |
| Housing stock/tenure | CBS StatLine | Gemeente | Yes |
| Housing prices | CBS (aggregate) / Kadaster (paid) | Gemeente (free) / Address (paid) |
| Employment/unemployment | UWV datasets | Arbeidsmarktregio (35) | Yes |
| Wages by sector | CBS CAO tables | Sector (national) | Yes |
| Crime/safety | CBS + Politie | Gemeente | Yes |
| Health/lifestyle | RIVM + CBS | Wijk (modelled) | Yes |
| Religion | CBS StatLine | Province | Yes |
| Media consumption | NMO/NOM | National (by demographics) | Summaries only |
| Social attitudes | LISS Panel | Province (individual-level) | Academic use |
| Election results | Kiesraad | Stembureau (polling station) | Yes |
| Social trust | ESS / Eurobarometer | National | Yes (GESIS) |

### Belgium -- Finest Granularity Available

| Category | Source | Finest Granularity | Open/Free |
|----------|--------|--------------------|-----------|
| Population demographics | StatBel | Statistical sector (~20,000) | Yes (CC BY 4.0) |
| Nationality/origin | StatBel | Commune (581) | Yes |
| Education level | Census 2021 | Commune | Yes |
| Fiscal income | StatBel open data | Statistical sector | Yes |
| Consumer spending | StatBel HBS | Region (3) | Aggregates free |
| Housing prices | StatBel real estate | Commune | Yes (to 2017 open; recent in reports) |
| Employment/unemployment | VDAB/FOREM/Actiris | Gemeente/Commune | Yes (per region) |
| Wages by occupation | StatBel | Province/Region | Yes |
| Crime | Federal Police | Commune / Police zone | Yes |
| Health | Sciensano HIS | Province/Region | Reports free, microdata by application |
| Religion | ESS / Eurobarometer | National (some NUTS2) | Yes (GESIS) |
| Media consumption | CIM | National (by demographics) | Summaries only |
| Social attitudes | RepResent / PartiRep | Individual-level | Academic application |
| Election results | SPF Interieur API | Commune | Yes (JSON API) |
| Social trust | ESS / Eurobarometer | National | Yes (GESIS) |
| Living conditions | EU-SILC (StatBel) | Region (3) | Microdata by application |

---

## Key Differences: NL vs BE for Persona Generation

1. **API maturity:** CBS (NL) has a proper OData REST API with no auth. StatBel (BE) relies on bulk file downloads -- no structured query API.
2. **Finest granularity:** NL wins with buurt-level data (~13,000 units). BE has statistical sector (~20,000) for population and fiscal income, but most other data is commune (581) or region (3) level.
3. **Income data:** NL has income at wijk/buurt level. BE has fiscal income at statistical sector level (excellent).
4. **Consumer spending:** Both weak -- national or regional only. No sub-regional consumer spending in either country.
5. **Linguistic complexity:** BE requires all data handling in two languages (NL and FR). Column names bilingual. Media landscape completely split.
6. **Election data:** BE has a proper JSON REST API with Swagger docs. NL uses EML/XML files (less convenient).
7. **Religion data:** NL publishes CBS tables with regional breakdown. BE has no official religion statistics (survey-based only).
8. **Media data:** Both proprietary at detailed level. NMO (NL) and CIM (BE) charge for audience data.
9. **Political parties:** BE has ~20+ parties split by language. NL has ~15 national parties. Both require extensive party mapping.
10. **Health data:** RIVM (NL) publishes modelled estimates at wijk level. Sciensano (BE) publishes at province/region level.
