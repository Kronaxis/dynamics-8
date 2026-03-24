# Canada: Comprehensive Data Source Catalogue for Synthetic Persona Generation

**Purpose:** Exhaustive reference of every actionable data source for generating census-weighted synthetic consumer personas for Canada, at the finest geographic granularity available.

**Last updated:** 2026-03-17

---

## 1. Census / Demographics

### 1.1 Statistics Canada Census of Population 2021

**Primary source for all demographic weighting.**

| Field | Detail |
|---|---|
| Agency | Statistics Canada (StatCan) |
| Coverage | Entire population of Canada (mandatory census) |
| Reference date | 11 May 2021 |
| Next census | 2026 |
| Licence | Open Government Licence - Canada (free, unrestricted use) |
| Formats | CSV, TAB, XML, JSON |

#### API Access: Three Distinct Services

**A. Web Data Service (WDS) -- for CANSIM/CODR tables**

| Field | Detail |
|---|---|
| Base URL | `https://www150.statcan.gc.ca/t1/wds/rest/` |
| Protocol | REST over HTTPS, returns JSON |
| Auth | None (free, no API key) |
| Rate limits | No published rate limit for standard queries |
| Documentation | https://www.statcan.gc.ca/en/developers/wds/user-guide |

Key methods (15 total):

| Method | HTTP | Purpose |
|---|---|---|
| `getAllCubesList` | GET | Full inventory of all data tables |
| `getAllCubesListLite` | GET | Lightweight table inventory |
| `getCubeMetadata` | POST | Metadata for a specific table |
| `getSeriesInfoFromCubePidCoord` | POST | Series info from product ID + coordinate |
| `getSeriesInfoFromVector` | POST | Series info from vector ID |
| `getDataFromCubePidCoordAndLatestNPeriods` | POST | Data by product ID, coordinate, N periods |
| `getDataFromVectorsAndLatestNPeriods` | POST | Data by vector ID, N periods |
| `getBulkVectorDataByRange` | POST | Bulk download by date range |
| `getDataFromVectorByReferencePeriodRange` | POST | Data by vector and reference period |
| `getChangedSeriesList` | POST | Changed series since date |
| `getChangedCubeList` | POST | Changed tables since date |
| `getChangedSeriesDataFromCubePidCoord` | POST | Changed data by product ID |
| `getChangedSeriesDataFromVector` | POST | Changed data by vector |
| `getFullTableDownloadCSV` | GET | Full table CSV download link |
| `getFullTableDownloadSDMX` | GET | Full table SDMX download link |

Example: `GET https://www150.statcan.gc.ca/t1/wds/rest/getFullTableDownloadCSV/98100006/en`

**B. Census Profile API (SDMX) -- for 2021 Census Profile data**

| Field | Detail |
|---|---|
| Base URL | `https://api.statcan.gc.ca/census-recensement/profile/sdmx/rest/` |
| Protocol | SDMX RESTful API |
| Auth | None |
| Formats | CSV, JSON, XML |
| Agency ID | `STC_CP` |
| Documentation | https://www12.statcan.gc.ca/wds-sdw/2021profile-profil2021-eng.cfm |

Supported resources: `datastructure`, `codelist`, `dataflow`. Dataflows are unique per geography level.

Key codelists:
- `CL_GEO_CMACA` -- Census Metropolitan Areas and Census Agglomerations
- `CL_GEO_PR` -- Provinces and Territories
- `CL_GEO_CD` -- Census Divisions
- `CL_GEO_CSD` -- Census Subdivisions
- `CL_GEO_CT` -- Census Tracts
- `CL_GEO_DA` -- Dissemination Areas

**C. Bulk Download -- for complete Census Profile CSVs**

| Field | Detail |
|---|---|
| URL | https://www12.statcan.gc.ca/census-recensement/2021/dp-pd/prof/details/download-telecharger.cfm |
| Format | CSV, TAB (compressed) |
| Scope | Complete Census Profile for all geographies within a hierarchy |
| Auth | None |
| Also on | Open Government Portal: https://open.canada.ca/data/en/dataset/750e6035-adf8-4426-966f-4c25b12a999e |

**Recommendation:** Bulk download is the best option for persona generation. Download the full Census Profile CSV for Census Divisions (293 areas), which includes 2,600+ characteristics per area. For finer granularity, download Census Tract or Dissemination Area files (larger but more precise).

#### Geographic Hierarchy (finest to coarsest)

| Level | Count | Typical Population | Notes |
|---|---|---|---|
| Dissemination Area (DA) | ~57,000 | 400-700 | Smallest standard unit. Full census data available. |
| Dissemination Block (DB) | ~490,000 | ~100 | Only population/dwelling counts. Too small for profiles. |
| Census Tract (CT) | ~6,600 | 2,500-8,000 | Only in CMAs/CAs with pop > 50,000. Best for urban personas. |
| Census Subdivision (CSD) | 5,162 | Varies widely | Municipalities. Many very small (<1,000). |
| Census Division (CD) | 293 | 20,000-500,000 | Regional districts/counties. **Best balance for persona generation.** |
| Census Metropolitan Area (CMA) | 41 | 100,000+ | Major urban centres. Cross-cuts CDs. |
| Census Agglomeration (CA) | 111 | 10,000-99,999 | Smaller urban centres. |
| Province/Territory (PT) | 13 | 40,000-14.8M | Too coarse for persona variation. |

**DGUID (Dissemination Geography Unique Identifier):** Standard identifier format linking all geography levels. Available in the Geographic Attribute File: https://open.canada.ca/data/en/dataset/1b3653d7-a48e-4001-8046-e6964bebe286

Boundary files (shapefiles): https://open.canada.ca/data/en/dataset/ef70dc3b-1069-4037-9bce-61f47e628a1d

#### Key Census Table Codes (98-10 Series)

| Dimension | Table | Geographic Levels | Notes |
|---|---|---|---|
| **Age/Sex** | 98-10-0002 | PT, CD, CSD, CMA, CA, CT | Age (single years) by sex. Released Feb 2022. |
| **Population counts** | 98-10-0006 | CMA, CA, Pop Centre | Population and land area. 100% data. |
| **Income** | 98-10-0055 | PT, CD, CMA, CA | Household total income groups by household characteristics. |
| **Income** | 98-10-0116 | PT, CD, CSD, CMA, CA | Individual income statistics by age, sex. |
| **Housing** | 98-10-0248 | PT, CMA, CA | Core housing need by tenure/mortgage/subsidy. |
| **Visible minority** | 98-10-0324 | PT, CMA, CA | Visible minority by generation status. |
| **Visible minority** | 98-10-0330 | PT | Visible minority by occupation, education, generation. |
| **Visible minority** | 98-10-0331 | PT, CMA, CA | Visible minority by income and generation. |
| **Visible minority** | 98-10-0342 | PT, CMA, CA | Religion by visible minority and generation. |
| **Visible minority** | 98-10-0347 | PT, CD, CMA, CA | Visible minority cross-tabs (general). |
| **Religion** | 98-10-0353-01 | PT | Religion by gender and age (province level). |
| **Religion** | 98-10-0353-02 | CMA, CA | Religion by gender and age (metro level). |
| **Education** | 98-10-0380 | PT, CD, CSD, CMA, CA | Highest certificate/diploma/degree. |
| **Education** | 98-10-0384 | PT, CMA, CA | Education by census year (time series). |
| **Education + Income** | 98-10-0410 | PT, CMA, CA | Income by education and major field of study. |
| **Education + Income** | 98-10-0411 | PT, CD, CSD | Income by education (census division level). |
| **Occupation** | 98-10-0451 | PT, CD, CMA, CA | Occupation by NOC (National Occupational Classification). |
| **Occupation + Income** | 98-10-0586 | PT | Employment income by occupation, visible minority, education. |
| **Language** | 98-10-0169 | PT, CMA, CA | Mother tongue by official language spoken, home language. |
| **Language** | 98-10-0364 | PT, CMA, CA | Mother tongue by generation status. |
| **Indigenous** | 98-10-0264 | PT | Indigenous identity by registered/treaty Indian status. |
| **Indigenous** | 98-10-0266 | PT, CD, CSD | Indigenous identity by registered status (CD level). |
| **Immigration** | 98-10-0307 | PT, CMA, CA | Immigrant status, place of birth, citizenship. |

All Census Profile downloads include all 2,600+ characteristics; individual tables offer more granular cross-tabulations.

#### Third Party API: CensusMapper / cancensus (R package)

| Field | Detail |
|---|---|
| API | https://censusmapper.ca/api |
| R package | `cancensus` (CRAN: https://cran.r-project.org/web/packages/cancensus/) |
| Python wrapper | `censusmapper` (limited) |
| Auth | Free API key from CensusMapper account signup |
| Coverage | Census years 1996, 2001, 2006, 2011, 2016, 2021 |
| Extra data | Annual taxfiler data at CT level (2000-2018): incomes, demographics |
| Granularity | All levels down to DA |
| Limitations | API quotas (generous but capped). Best for programmatic access. |

Key functions:
- `list_census_datasets()` -- list available census years
- `list_census_regions(dataset)` -- all named regions and IDs
- `search_census_vectors(search_term, dataset)` -- find variable codes
- `get_census(dataset, regions, vectors, level)` -- retrieve data
- `get_census_geometry(...)` -- retrieve with spatial boundaries

#### Indigenous Population Data

| Source | Detail |
|---|---|
| Indigenous Population Profile | https://www12.statcan.gc.ca/census-recensement/2021/dp-pd/ipp-ppa/index.cfm |
| Table 98-10-0264 | Indigenous identity by Registered/Treaty Indian status, by Indigenous geography |
| Table 98-10-0266 | Indigenous identity by CD and CSD |
| Geographic levels | PT, CMA, CA, CD, CSD, Metis settlements, Inuit regions, Indian band/Tribal Council areas |
| Total count | 1,807,250 Indigenous people (5% of population): First Nations 1,048,405; Metis 624,220; Inuit 70,545 |

#### Language Data

| Source | Detail |
|---|---|
| Table 98-10-0169 | Mother tongue by official language spoken, home language |
| Table 98-10-0364 | Mother tongue by generation status and languages known |
| Census Profile | Knowledge of official languages, mother tongue, home language, language at work |
| Key stat | 9M people had mother tongue other than English/French. 1 in 8 speak non-official language at home. |
| Geographic levels | PT, CMA, CA, CSD, CT |

---

## 2. Consumer Spending

### 2.1 Survey of Household Spending (SHS)

| Field | Detail |
|---|---|
| Agency | Statistics Canada |
| Frequency | Annual |
| Latest | 2023 (released May 2025) |
| Licence | Open Government Licence - Canada |

#### Key Tables

| Table | Title | Geography | Categories |
|---|---|---|---|
| **11-10-0222-01** | Household spending, Canada, regions and provinces | Canada, regions, provinces | 300+ spending categories |
| **11-10-0197-01** | Household spending, summary-level categories, by province, territory and selected metropolitan areas | PT + selected CMAs | 19 summary categories |
| **11-10-0224-01** | Household spending by household type | Canada, provinces | By household composition |
| **36-10-0225-01** | Detailed household final consumption, provincial/territorial | PT | National accounts basis |

Summary categories in 11-10-0197: Total expenditure, Total current consumption, Food, Shelter, Household operation, Household furnishings/equipment, Clothing/accessories, Transportation, Health care, Personal care, Recreation, Education, Reading materials, Tobacco/alcohol, Games of chance, Miscellaneous, Personal insurance/pension contributions, Gifts, Security.

**Geographic granularity:** Province/territory level for most tables. Selected CMAs in 11-10-0197 only (typically ~15 largest metros). **No CD or CT level spending data.**

**Limitation:** CV (coefficient of variation) >= 35% suppressed as unreliable. Metro-level estimates often suppressed for smaller CMAs.

**API access:** Via WDS: `POST https://www150.statcan.gc.ca/t1/wds/rest/getFullTableDownloadCSV/11100222/en`

**Open Government Portal:** https://open.canada.ca/data/en/dataset/9f2d519e-4b2c-4fa7-af79-032142b43760

---

## 3. Media Consumption

### 3.1 Reuters Institute Digital News Report -- Canada

| Field | Detail |
|---|---|
| URL | https://reutersinstitute.politics.ox.ac.uk/digital-news-report/2024/canada |
| Frequency | Annual (published June) |
| Format | PDF report + interactive data explorer |
| Licence | Free to access; academic use permitted |
| Coverage | National sample (~2,000 respondents) |
| Variables | News sources used, trust in news, platform usage, social media for news, willingness to pay, podcast consumption, news avoidance |
| Granularity | National only (no provincial breakdown) |
| Key finding (2024) | News blocked on Meta platforms (Online News Act). Google paying ~US$73M/year. Village Media expanding local digital model in Ontario. |
| Limitation | No raw data download. Summary tables and charts only. Use for aggregate calibration, not geographic weighting. |

### 3.2 CRTC Communications Market Reports

| Field | Detail |
|---|---|
| Agency | Canadian Radio-television and Telecommunications Commission |
| URL | https://crtc.gc.ca/eng/publications/reports/PolicyMonitoring/ |
| Open Data | https://crtc.gc.ca/eng/publications/reports/policymonitoring/cmrd.htm |
| Frequency | Annual report + quarterly data updates |
| Latest data | 2025-Q3 (quarterly), January 2026 (monthly) |
| Format | 400+ open datasets (CSV/Excel) |
| Licence | Open Government Licence - Canada |
| Coverage | National and by province/territory |
| Variables | TV tuning hours, radio tuning hours, streaming hours, internet subscriptions, mobile subscriptions, broadband penetration, cord-cutting rates, OTT revenue |
| Key stats (latest) | Radio: 1.16B hours/quarter (-0.9%). TV: 7.9B hours/quarter (-1.8%). AV streaming: 2.83B hours/quarter (+13.4%). |
| Granularity | National and provincial. Some metrics by CMA. |

### 3.3 Social Media Usage

| Source | Detail |
|---|---|
| Statistics Canada (Canadians' assessments of social media) | https://www150.statcan.gc.ca/n1/pub/36-28-0001/2021003/article/00004-eng.htm |
| DataReportal Digital 2025 Canada | https://datareportal.com/reports/digital-2025-canada |
| StatCounter Social Media Stats Canada | https://gs.statcounter.com/social-media-stats/all/canada |
| Licence | DataReportal: free summary; full data requires subscription. StatCounter: free for aggregate. StatCan: OGL. |
| Key stats | 95.2% internet penetration. 79.4% social media penetration (31.7M users). 6h45m daily internet. 1h53m daily social media. |
| Platform share | Facebook (~30M users), Instagram (~20M), TikTok, X/Twitter, LinkedIn, Snapchat, Reddit |
| Age patterns | 90% of under-34s are regular social media users; 70% of 50-64; 50% of 65+ |
| Granularity | National only. Age/gender breakdowns available. No provincial or sub-provincial. |
| Limitation | No Canadian-specific survey with provincial social media breakdowns in public domain. Use national averages + age/urban-rural adjustments. |

### 3.4 Local Newspaper Coverage by Province/City

Canada's newspaper landscape is highly consolidated. Major chains:

| Chain | Reach | Key Papers |
|---|---|---|
| Postmedia | National (English) | National Post, Vancouver Sun, Calgary Herald, Edmonton Journal, Ottawa Citizen, Montreal Gazette, Toronto Sun, Winnipeg Sun |
| Torstar | Ontario | Toronto Star, Hamilton Spectator |
| Globe and Mail | National | Globe and Mail (independent) |
| Quebecor (TVA/Le Journal) | Quebec (French) | Le Journal de Montreal, Le Journal de Quebec, TVA Nouvelles |
| La Presse | Quebec (French) | La Presse (digital-only since 2018) |
| Le Devoir | Quebec (French) | Le Devoir (independent) |
| SaltWire | Atlantic | Chronicle Herald (Halifax), Cape Breton Post, The Guardian (PEI) |
| Black Press (now Canadian-owned) | BC, Alberta | 80+ community papers |
| Village Media | Ontario | 40+ local digital news sites |
| CBC/Radio-Canada | National (both languages) | Public broadcaster, digital + TV + radio |

---

## 4. Social Attitudes / Institutional Trust

### 4.1 Canadian Election Study (CES)

| Field | Detail |
|---|---|
| URL | https://ces-eec.arts.ubc.ca/ |
| Data repository | Harvard Dataverse: https://dataverse.harvard.edu/dataset.xhtml?persistentId=doi:10.7910/DVN/XBZHKC (2021) |
| Also on | Borealis Data (Canadian institutional repository) |
| Coverage | 2021 CES (web), 2019 CES (web n=37,822 + phone n=4,021), 2015, and back to 1965 |
| Format | SPSS, Stata, CSV |
| Licence | Free for academic/research use (requires Dataverse account) |
| Variables | Vote choice, party identification, leader evaluations, issue positions, values, ideology (left-right), economic perceptions, immigration attitudes, trust in institutions, democratic satisfaction, media consumption, identity, demographics |
| R packages | `cesR` (https://hodgettsp.github.io/cesR/) and `ces` (CRAN) |
| Granularity | Individual-level (respondent). Province available as variable. No sub-provincial geography (respondent privacy). |
| Update frequency | Each federal election |
| **Strength for persona generation:** Best Canadian source for political attitudes, vote intention, issue positions, ideology, and party affiliation by demographic cross-tabs. |

### 4.2 Environics Institute -- Focus Canada

| Field | Detail |
|---|---|
| URL | https://www.environicsinstitute.org/ |
| Project listing | https://www.environicsinstitute.org/projects/listing/-in-tags/type/focus-canada |
| Coverage | Continuous since 1977 (originally Environics Research Group syndicated survey). Published free since ~2010. |
| Format | PDF reports (not raw data). Summary tables, charts, and crosstabs in reports. |
| Licence | Free to read. Raw data not publicly downloadable. |
| Topics | National identity, immigration attitudes, multiculturalism, social values, Indigenous reconciliation, regional identity, trust in institutions, environment, healthcare, economic outlook |
| Granularity | National. Regional breakdowns (West, Ontario, Quebec, Atlantic) in some reports. |
| Key reports | Trust in Political Institutions (2025, 2024), Focus Canada (annual), Canadian Confederation Study, Diversity & Inclusion, Truth and Reconciliation |
| Limitation | No machine-readable data download. Must extract values from PDF reports manually. Useful for attitude calibration but not direct weighting. |

### 4.3 ISSP (International Social Survey Programme) -- Canada

| Field | Detail |
|---|---|
| URL | https://www.gesis.org/en/issp/data-and-documentation |
| Data access | GESIS Data Catalogue (free download after registration) |
| Format | SPSS, Stata |
| Licence | Free for academic/research use |
| Canada participation | Intermittent. Participated in: Social Inequality, Citizenship, Work Orientations, Family and Gender Roles, National Identity (various years). **Not a member for every module.** |
| Variables | Module-dependent: attitudes to government, inequality, religion, environment, health, citizenship, work, family |
| Granularity | Individual-level. Country only (no sub-national geography). |
| Limitation | Canada's participation is sporadic. Check module-by-module availability. Useful for cross-national attitude benchmarking. |

### 4.4 Edelman Trust Barometer -- Canada

| Field | Detail |
|---|---|
| 2024 report | https://www.edelman.ca/sites/g/files/aatuss376/files/2024-03/2024%20Edelman%20Trust%20Barometer_Canada%20Report_EN_0.pdf |
| 2025 report | https://www.edelman.ca/sites/g/files/aatuss376/files/2025-03/2025%20Edelman%20Trust%20Barometer_Canada%20Report_MASTER.pdf |
| Archive | https://www.edelman.com/trust/archive |
| Format | PDF (free download) |
| Licence | Free to access |
| Coverage | ~1,150 Canadian respondents per year |
| Variables | Trust in business, government, media, NGOs. Trust in CEO, journalist, government official. Innovation optimism/fear. Information sources. |
| Key finding (2024) | Business most trusted institution in Canada (57%). Government in distrust. |
| Granularity | National only. Some demographic cross-tabs (age, income, gender). No provincial. |
| Limitation | PDF only. No raw data. Summary-level trust scores for calibration. |

### 4.5 Angus Reid Institute

| Field | Detail |
|---|---|
| URL | https://angusreid.org/ |
| Format | PDF reports with summary data tables and charts |
| Licence | Free to access reports. Raw data not publicly available. |
| Topics | Federal/provincial politics, trust, religion, social issues, economy, environment, housing, immigration, healthcare, Indigenous issues |
| Key surveys | Faith in Canada (religion spectrum: Non-Believers, Spiritually Uncertain, Privately Faithful, Religiously Committed), Trust and Politics, Election tracking, Housing affordability |
| Granularity | National. Some regional breakdowns (West, Ontario, Quebec, Atlantic). |
| Update frequency | Multiple reports per month |
| Limitation | No machine-readable data download. Proprietary micro-data. Good for attitude calibration from summary tables. |

---

## 5. Political Landscape

### 5.1 Elections Canada -- Federal Election Results

| Field | Detail |
|---|---|
| URL | https://www.elections.ca/content.aspx?section=res&dir=rep/off&document=index&lang=e |
| Open Data Portal | https://open.canada.ca/data/en/organization/elections?res_type=dataset |
| Format | CSV (compressed ZIP) |
| Licence | Open Government Licence - Canada |
| Auth | None |

#### Available Elections (FPTP, 338 ridings since 2015)

| Election | Year | Portal Link |
|---|---|---|
| 44th General Election | 2021 | https://open.canada.ca/data/en/dataset/065439a9-c194-4259-9c95-245a852be4a1 |
| 43rd General Election | 2019 | https://open.canada.ca/data/en/dataset/199e5070-2fd5-49d3-aa21-aece08964d18 |
| 42nd General Election | 2015 | https://open.canada.ca/data/en/dataset/775f3136-1aa3-4854-a51e-1a2dab362525 |

Data fields: Electoral district number, electoral district name, province, candidate name, party affiliation, votes obtained, percentage of votes, elected/defeated, total valid votes, rejected ballots, total ballots cast, electors on list, voter turnout.

**Granularity:** Poll-by-poll results within each of 338 federal electoral districts. This is the finest level: individual polling stations.

**Electoral district boundaries:** https://open.canada.ca/data/en/dataset/47a0f098-7445-41bb-a147-41686b692887

### 5.2 Provincial Election Results

Each province maintains its own electoral authority. Key sources:

| Province | Authority | Data Availability |
|---|---|---|
| Ontario | Elections Ontario (elections.on.ca) | Results explorer with CSV export. Historical back to 1867. |
| Quebec | Elections Quebec (electionsquebec.qc.ca) | Open data with riding-level results. |
| British Columbia | Elections BC (elections.bc.ca) | Historical results, open data. |
| Alberta | Elections Alberta (elections.ab.ca) | Results by electoral division. |
| Saskatchewan | Elections Saskatchewan | Results by constituency. |
| Manitoba | Elections Manitoba | Results by electoral division. |
| Nova Scotia | Elections Nova Scotia | Results by electoral district. |
| New Brunswick | Elections New Brunswick | Results by electoral district. |
| PEI | Elections PEI | Results by district. |
| Newfoundland & Labrador | Elections NL | Results by district. |
| Territories | Respective electoral offices | Smaller datasets. |

**Limitation:** No single unified dataset for all provincial elections. Each must be scraped or downloaded individually. Formats vary (CSV, PDF, HTML tables).

### 5.3 Polling Aggregation -- 338Canada

| Field | Detail |
|---|---|
| URL | https://338canada.com/ |
| Author | Philippe J. Fournier |
| Coverage | Federal + all provinces |
| Update | Federal: every Sunday (outside campaigns). Campaigns: daily. |
| Data | Riding-level projections (338 federal + provincial ridings), poll aggregation, seat projections, popular vote estimates |
| Format | Web-based visualisations. No formal API or CSV download. |
| Licence | Proprietary. Scraping required for data extraction. |
| Limitation | Not an official data source. Model-based projections, not raw results. Useful for current political landscape context but not for historical weighting. |

### 5.4 Harvard Dataverse -- Canadian Elections Datasets

| Field | Detail |
|---|---|
| URL | https://semrasevi.com/candidates-in-canadian-elections-datasets/ |
| Coverage | 46,526+ federal candidates + 15,500+ Ontario provincial candidates |
| Format | CSV, Stata |
| Licence | Academic use |
| Variables | Candidate name, party, riding, votes, elected status, gender, incumbency |

---

## 6. Employment / Salary

### 6.1 Statistics Canada Labour Force Survey (LFS)

| Field | Detail |
|---|---|
| Agency | Statistics Canada |
| Frequency | Monthly |
| Sample | ~56,000 households per month |
| Licence | Open Government Licence - Canada |

#### Key Tables

| Table | Title | Geography | Variables |
|---|---|---|---|
| **14-10-0417-01** | Employee wages by occupation, annual | Canada, PT, CMA, Economic Region | Average/median hourly and weekly wages by NOC, type of work, gender, age |
| **14-10-0340-01** | Employee wages by occupation (1997-2022, inactive) | Canada, PT | Historical wage series |
| **14-10-0416-01** | Labour force characteristics by occupation, annual | Canada, PT | Employment, unemployment by NOC |
| **14-10-0287-01** | Labour force characteristics, monthly (seasonally adjusted) | Canada, PT, CMA | Headline employment/unemployment rates |
| **14-10-0064-01** | Employee wages by industry, annual | Canada, PT | By NAICS industry |

**Geographic granularity:** National, provincial, CMA (41 metro areas), Economic Region (~76 regions). Some tables go to CMA; others only provincial.

### 6.2 Census 2021 Occupation and Income Data

| Table | Title | Geography |
|---|---|---|
| **98-10-0451** | Occupation by NOC (2021) | PT, CD, CMA, CA |
| **98-10-0411** | Employment income by education | PT, CD, CSD |
| **98-10-0586** | Employment income by occupation, visible minority, education | PT |
| **98-10-0412** | Employment income by occupation and field of study | Canada |

### 6.3 National Occupational Classification (NOC) 2021

| Field | Detail |
|---|---|
| URL | https://www150.statcan.gc.ca/n1/en/catalogue/12-583-X |
| Structure | 5-digit hierarchy: 10 broad categories, 45 major groups, 89 minor groups, 162 unit groups, 516 occupations |
| NOC-to-SOC crosswalk | Partial. LMIC report no. 35 (https://lmic-cimt.ca/publications-all/lmi-insight-report-no-35/) provides concordance, but 62 of 516 NOC codes (7%) have no clear SOC equivalent (~10% of workforce). |
| O*NET mapping | Via The Dais crosswalk (https://dais.ca/reports/crosswalk-blog-post/). Updated versions cover NOC 2001 onwards to O*NET-SOC. |

### 6.4 Canada Job Bank Wage Data

| Field | Detail |
|---|---|
| URL | https://www.jobbank.gc.ca/trend-analysis/search-wages |
| Open Data (CSV) | https://open.canada.ca/data/en/dataset/adad580f-76b0-4502-bd05-20c125de9116 |
| Coverage | 516 NOC 2021 occupations |
| Geography | National, provincial/territorial, economic region |
| Variables | Low, median, high wages per occupation per region |
| Update | Annual (typically autumn) |
| Format | CSV download from Open Government Portal |
| Methodology | Joint ESDC / Statistics Canada methodology using LFS, Census, and administrative data |
| Auth | None |
| Licence | Open Government Licence - Canada |

**Recommendation:** Job Bank wage CSV is the most accessible programmatic source for occupation-specific wages by region. Complement with Census table 98-10-0411 for education-income cross-tabs at CD level.

---

## 7. Housing Market

### 7.1 CMHC (Canada Mortgage and Housing Corporation)

#### Rental Market Survey

| Field | Detail |
|---|---|
| URL | https://www.cmhc-schl.gc.ca/professionals/housing-markets-data-and-research/housing-data/data-tables/rental-market/rental-market-report-data-tables |
| Frequency | Annual (published December) |
| Latest | 2025 (released 11 December 2025) |
| Format | Microsoft Excel (.xlsx) |
| Licence | Open |
| Variables | Vacancy rates, average rents (by bedroom count), turnover rates, universe counts |
| Geography | Canada, all provinces, major centres (~40 CMAs). 17 centres include secondary market (condo rental). |
| Auth | None |

#### Housing Market Information Portal (HMIP)

| Field | Detail |
|---|---|
| URL | https://www.cmhc-schl.gc.ca/hmiportal |
| Interactive tool | https://www03.cmhc-schl.gc.ca/hmip-pimh/en/TableMapChart |
| Variables | Starts, completions, under construction, absorption, price, vacancy, rent |
| Geography | National to neighbourhood level |
| Format | Interactive (CSV export available) |
| R package | `cmhc` (https://mountainmath.github.io/cmhc/) -- API wrapper for HMIP portal data |
| Open Data | https://open.canada.ca/data/en/dataset/c2a1fdbf-d9b7-4c84-b7eb-c845b6ffd5e6 |

#### StatCan Table on CMHC Rents

| Table | Title |
|---|---|
| **34-10-0133-01** | CMHC average rents for areas with population 10,000+ |

### 7.2 Census 2021 Housing Data

| Table | Title | Geography |
|---|---|---|
| **98-10-0248** | Core housing need by tenure/mortgage/subsidy | PT, CMA, CA |
| Census Profile | Dwelling type, tenure, shelter cost, housing suitability, condition | All levels to DA |

### 7.3 CREA (Canadian Real Estate Association)

| Field | Detail |
|---|---|
| URL | https://www.crea.ca/housing-market-stats/ |
| MLS HPI Tool | https://www.crea.ca/housing-market-stats/mls-home-price-index/hpi-tool/ |
| National Price Map | https://www.crea.ca/housing-market-stats/canadian-housing-market-stats/national-price-map/ |
| Data download | MLS HPI data and Quarterly Forecasts available as Excel (.xlsx) |
| Access | Public dashboard and map; some detailed data restricted to REALTORS |
| Variables | Benchmark prices by dwelling type, sales volumes, new listings, months of inventory, average prices |
| Geography | Major markets and provinces (monthly) |
| Licence | Public summary data freely accessible. Detailed MLS data restricted. |
| Limitation | Not all data publicly downloadable. Best used for benchmark home prices per city. |

---

## 8. Crime / Safety

### 8.1 Uniform Crime Reporting (UCR) Survey

| Field | Detail |
|---|---|
| Agency | Statistics Canada |
| Survey | UCR (census of all police-reported crime) |
| Licence | Open Government Licence - Canada |
| History | Annual since 1962 |

#### Key Tables

| Table | Title | Geography |
|---|---|---|
| **35-10-0026-01** | Crime Severity Index and weighted clearance rates | Canada, PT, CMA |
| **35-10-0177-01** | Incident-based crime statistics by detailed violations | Canada, PT, CMA, Canadian Forces Military Police |

Table 35-10-0177 provides incident counts for all Criminal Code offences, broken down by detailed violation type, for every CMA (~41 areas) plus provinces and territories.

Crime Severity Index: Weighted measure where more serious crimes have higher weights. Both overall CSI and violent CSI available.

**Geographic granularity:** Provincial and CMA (41 metro areas). Not available at CD, CSD, or CT level from StatCan. Some municipal police services publish neighbourhood-level crime data independently (e.g., Toronto Police open data).

**Update frequency:** Annual (typically July release for prior year).

### 8.2 General Social Survey -- Canadians' Safety (Victimization)

| Field | Detail |
|---|---|
| Survey | GSS Cycle 34 (2019 is latest) |
| PUMF | Public Use Microdata File available (provincial geography) |
| Catalogue | 12M0026X (2014 cycle), newer via StatCan Research Data Centres |
| Variables | Self-reported victimization (violent, household, sexual), reporting to police, perceptions of safety, neighbourhood disorder, confidence in justice system |
| Geography | Provincial in PUMF. Health region in master file (restricted). |
| Table 35-10-0164-01 | Violent victimization by reporting to police, gender, geography |
| Licence | OGL for PUMF. Master file requires Research Data Centre access. |
| Limitation | Self-reported (not police-reported). 5-year cycle. 2019 is latest available. |

---

## 9. Religion

### 9.1 Census 2021 Religion Data

| Field | Detail |
|---|---|
| Tables | 98-10-0353-01 (PT), 98-10-0353-02 (CMA/CA), 98-10-0353-03 (CMA/CA by PT) |
| Cross-tab | 98-10-0342 -- Religion by visible minority and generation status (PT, CMA, CA) |
| Categories | 25 religion categories (including subcategories of Christianity, Islam, Hinduism, Sikhism, Buddhism, Judaism, Indigenous spirituality, no religion, etc.) |
| Sample | 25% of private households |
| Geographic levels | PT, CMA, CA (not below CMA) |
| Release | October 2022 |

**Key statistics (2021 Census):**
- Christian: 53.3% (19.3M) -- declining from 67.3% in 2011
- No religion: 34.6% (12.6M) -- rising from 23.9% in 2011
- Muslim: 4.9% (1.8M)
- Hindu: 2.3%
- Sikh: 2.1%
- Buddhist: 1.0%
- Jewish: 0.9%
- Indigenous spirituality: 0.2%
- Other: 0.7%

### 9.2 Angus Reid Institute -- Religion Spectrum

| Field | Detail |
|---|---|
| Report | "Canada across the religious spectrum" (2022) |
| URL | https://angusreid.org/wp-content/uploads/2022/04/2022.04.18_Religious_Spectrum.pdf |
| Framework | Four-category spectrum: Non-Believers, Spiritually Uncertain, Privately Faithful, Religiously Committed |
| Data | 63% describe as spiritual. 39% "spiritual but not religious." 24% "spiritual and religious." |
| Regional breakdown | Some provincial/regional breakdowns in report |
| Limitation | PDF summary only. No raw data. |

---

## 10. Health

### 10.1 Canadian Community Health Survey (CCHS)

| Field | Detail |
|---|---|
| Agency | Statistics Canada |
| Survey | CCHS Annual Component (SDDS 3226) |
| Coverage | 10 provinces (excludes territories), population aged 12+, ~65,000 respondents/year |
| Licence | OGL for published tables. PUMF for microdata (registration required). |
| Frequency | Annual (with 2-year combined estimates for smaller areas) |

#### Key Tables

| Table | Title | Geography |
|---|---|---|
| **13-10-0096-01** | Health characteristics, annual estimates | Canada, PT |
| **13-10-0113-01** | Health characteristics, 2-year estimates | PT, Health Region |
| **13-10-0905-01** | Health indicator statistics, annual | PT |
| **13-10-0906-01** | Health indicators by income quintile and education | PT |
| **13-10-0601-01** | CCHS indicator profile (older) | PT, CMA |

Variables: Self-rated health, chronic conditions (asthma, diabetes, heart disease, cancer, arthritis, mood disorders, anxiety, COPD), BMI, physical activity, smoking, alcohol use, fruit/vegetable consumption, influenza vaccination, contact with health professionals, unmet health care needs, mental health (self-rated), life satisfaction, sense of belonging.

**Geographic granularity:** Provincial (annual estimates). Health Region (~100 regions) for 2-year combined estimates. CMA in some cross-tabs.

### 10.2 CCHS -- Mental Health Module

| Field | Detail |
|---|---|
| Survey | CCHS -- Mental Health (SDDS 5015) |
| Latest | 2012 (dedicated module); annual CCHS includes limited mental health items |
| Coverage | 10 provinces, population 15+ |
| Variables | Prevalence of mood disorders (depression, mania), anxiety disorders (panic, agoraphobia, social phobia), substance use disorders, suicidal ideation, mental health service use, disability days |
| Geography | Provincial (in PUMF). Health region in master file (restricted). |
| Mental Health and Access to Care Survey (MHACS, 2022) | https://www23.statcan.gc.ca/imdb/p2SV.pl?Function=getSurvey&SDDS=5015 -- newer replacement |

### 10.3 CIHI (Canadian Institute for Health Information)

| Field | Detail |
|---|---|
| URL | https://www.cihi.ca/ |
| Wait times tool | https://www.cihi.ca/en/explore-wait-times-for-priority-procedures-across-canada |
| Data download | XLSX data tables for wait times (2025 latest) |
| Variables | Wait times (median, 90th percentile) for hip replacement, knee replacement, cataract surgery, radiation therapy, hip fracture repair, cancer surgeries, CT/MRI scans, CABG |
| Geography | National, provincial (6 years). Health region for hip/knee replacement. |
| Other indicators | Perceived health, healthcare spending, hospital beds, physician supply, C-section rates |
| Licence | CIHI open data (free download for published indicators) |
| Limitation | Not raw microdata. Aggregate indicators only. Useful for healthcare system calibration by province. |
| Mental health wait times | https://www.cihi.ca/en/indicators/wait-times-for-community-mental-health-counselling |

---

## 11. News RSS Feeds

### 11.1 National / Major English-Language Feeds

| Outlet | RSS URL | Scope |
|---|---|---|
| **CBC News -- Top Stories** | `https://www.cbc.ca/cmlink/rss-topstories` | National |
| **CBC News -- Canada** | `https://rss.cbc.ca/lineup/canada.xml` | National |
| **CBC News -- World** | `https://rss.cbc.ca/lineup/world.xml` | International |
| **CBC News -- Politics** | `https://rss.cbc.ca/lineup/politics.xml` | Federal politics |
| **CBC News -- Business** | `https://rss.cbc.ca/lineup/business.xml` | Business |
| **CBC News -- Technology** | `https://rss.cbc.ca/lineup/technology.xml` | Technology |
| **CTV News -- Top Stories** | `https://www.ctvnews.ca/rss/ctvnews-ca-top-stories-public-rss-1.822009` | National |
| **Global News** | `https://globalnews.ca/feed/` | National |
| **National Post** | `https://nationalpost.com/feed/` | National (centre-right) |
| **Globe and Mail** | `https://www.theglobeandmail.com/arc/outboundfeeds/rss/category/canada/` | National (centre) |
| **Toronto Star** | `https://www.thestar.com/search/?f=rss&t=article&c=news*` | National/Ontario |
| **Maclean's** | `https://macleans.ca/feed/` | National magazine |
| **Financial Post** | `https://financialpost.com/feed/` | Business/finance |
| **The Narwhal** | `https://thenarwhal.ca/feed/` | Investigative/environment |
| **The Conversation (Canada)** | `https://theconversation.com/ca/articles.atom` | Academic/analysis |

Full CBC RSS index: https://www.cbc.ca/rss/

### 11.2 French-Language (Quebec) Feeds

| Outlet | RSS URL | Scope |
|---|---|---|
| **Radio-Canada -- Nouvelles** | `https://ici.radio-canada.ca/rss/4159` | National (French) |
| **Radio-Canada -- Politique** | `https://ici.radio-canada.ca/rss/4175` | Federal/Quebec politics |
| **La Presse -- Actualites** | `https://www.lapresse.ca/actualites/rss` | Quebec/National |
| **La Presse -- Manchettes** | `https://www.lapresse.ca/actualites/manchettes/rss` | Headlines only |
| **Le Devoir** | `https://www.ledevoir.com/rss/manchettes.xml` | Quebec (independent) |
| **Le Journal de Montreal** | Via itesmedia: sorted by category (see https://support.itesmedia.tv/rss-feed-1) | Quebec (tabloid) |
| **TVA Nouvelles** | `https://www.tvanouvelles.ca/rss.xml` | Quebec/National TV news |

Full Radio-Canada RSS index: https://ici.radio-canada.ca/info/rss
Full Le Devoir RSS index: https://www.ledevoir.com/flux-rss

### 11.3 Provincial / Regional Feeds

| Region | Outlet | RSS URL |
|---|---|---|
| **British Columbia** | Vancouver Sun | `https://vancouversun.com/feed/` |
| **British Columbia** | The Province | `https://theprovince.com/feed/` |
| **British Columbia** | CBC BC | `https://rss.cbc.ca/lineup/canada-britishcolumbia.xml` |
| **Alberta** | Calgary Herald | `https://calgaryherald.com/feed/` |
| **Alberta** | Edmonton Journal | `https://edmontonjournal.com/feed/` |
| **Alberta** | CBC Calgary | `https://rss.cbc.ca/lineup/canada-calgary.xml` |
| **Alberta** | CBC Edmonton | `https://rss.cbc.ca/lineup/canada-edmonton.xml` |
| **Saskatchewan** | CBC Saskatchewan | `https://rss.cbc.ca/lineup/canada-saskatchewan.xml` |
| **Manitoba** | CBC Manitoba | `https://rss.cbc.ca/lineup/canada-manitoba.xml` |
| **Manitoba** | Winnipeg Free Press | `https://www.winnipegfreepress.com/rss/` |
| **Ontario** | Ottawa Citizen | `https://ottawacitizen.com/feed/` |
| **Ontario** | CBC Toronto | `https://rss.cbc.ca/lineup/canada-toronto.xml` |
| **Ontario** | CBC Ottawa | `https://rss.cbc.ca/lineup/canada-ottawa.xml` |
| **Atlantic** | CBC Nova Scotia | `https://rss.cbc.ca/lineup/canada-novascotia.xml` |
| **Atlantic** | CBC New Brunswick | `https://rss.cbc.ca/lineup/canada-newbrunswick.xml` |
| **Atlantic** | CBC PEI | `https://rss.cbc.ca/lineup/canada-pei.xml` |
| **Atlantic** | CBC NL | `https://rss.cbc.ca/lineup/canada-newfoundland.xml` |
| **North** | CBC North | `https://rss.cbc.ca/lineup/canada-north.xml` |
| **Quebec** | CBC Montreal | `https://rss.cbc.ca/lineup/canada-montreal.xml` |
| **Quebec** | Montreal Gazette | `https://montrealgazette.com/feed/` |

**Curated OPML file:** https://github.com/plenaryapp/awesome-rss-feeds/blob/master/countries/without_category/Canada.opml

---

## Summary: Recommended Data Pipeline for Canadian Persona Generation

### Phase 1: Census Skeleton (Demographics)

**Source:** Statistics Canada Census Profile 2021 (bulk CSV download)
**Geography:** Census Division (293 units) -- best balance of granularity and statistical reliability
**Method:** Download complete Census Profile CSVs. Extract distributions for:
- Age/sex (single years)
- Visible minority categories
- Indigenous identity (First Nations, Metis, Inuit)
- Education (highest certificate/diploma/degree)
- Individual income bands
- Household income bands
- Occupation (NOC major groups)
- Mother tongue / official language knowledge
- Religion
- Housing tenure (own/rent), dwelling type
- Immigration status, period of immigration
- Commute mode, commute duration

### Phase 2: Economic Overlay

**Sources:**
- Job Bank wage CSV (occupation x province/economic region) -- median/low/high wages
- LFS Table 14-10-0417 (wages by occupation, CMA level)
- SHS Table 11-10-0222 (spending patterns, provincial)
- CMHC Rental Market Survey (rent levels by CMA)
- Census Profile (household income, shelter cost)

### Phase 3: Political/Attitudinal Layer

**Sources:**
- Elections Canada CSV (riding-level results, 3 elections) -- map ridings to CDs
- Canadian Election Study microdata (attitude distributions by demographics)
- Environics Focus Canada (regional attitude calibration)
- Edelman Trust Barometer (institutional trust levels)
- CCHS (health behaviours, self-rated health by province)
- Census religion data (98-10-0353)

### Phase 4: Media/Digital Layer

**Sources:**
- CRTC Communications Market Reports (TV/radio/streaming hours)
- Reuters Digital News Report (news platform usage)
- DataReportal/StatCounter (social media platform penetration by age)
- Local newspaper mapping (province-to-outlet lookup table)
- RSS feeds (populate media diet per persona)

### Phase 5: Validation

**Cross-reference:** Census-weighted distributions against published StatCan totals for each CD. Enforce diversity rules (ethnicity, age, gender, religion distributions must fall within confidence intervals of census data for the assigned geography).

---

## Licensing Summary

| Source | Licence | Machine-Readable Download | Auth Required |
|---|---|---|---|
| Statistics Canada (all tables/census) | Open Government Licence - Canada | Yes (CSV, JSON via WDS) | No |
| CensusMapper/cancensus | CC-compatible | Yes (R/Python API) | Free API key |
| Elections Canada | Open Government Licence - Canada | Yes (CSV) | No |
| CMHC | Open | Yes (Excel, via R package) | No |
| CRTC | Open Government Licence - Canada | Yes (CSV/Excel) | No |
| Job Bank wages | Open Government Licence - Canada | Yes (CSV) | No |
| CIHI | CIHI open data | Yes (Excel) | No |
| Canadian Election Study | Academic/research | Yes (SPSS/Stata/CSV) | Dataverse account |
| ISSP | Academic/research | Yes (SPSS/Stata) | GESIS account |
| Environics Institute | Free to read | No (PDF reports only) | No |
| Angus Reid Institute | Free to read | No (PDF reports only) | No |
| Edelman Trust Barometer | Free to read | No (PDF reports only) | No |
| Reuters Digital News Report | Free to read | No (interactive tool only) | No |
| CREA housing prices | Partially public | Partial (public Excel for HPI) | No for public data |
| 338Canada | Proprietary | No | N/A |

---

## Key Differences from UK Data Landscape

| Dimension | UK | Canada |
|---|---|---|
| Census agency | ONS (England/Wales), NRS (Scotland), NISRA (NI) | Statistics Canada (unified) |
| Finest standard unit | Output Area (~300 people) | Dissemination Area (~400-700 people) |
| Electoral geography | 650 parliamentary constituencies | 338 federal electoral districts (ridings) |
| Persona generation target | Constituency (650 x 100 = 65k) | Census Division (293) or Federal Riding (338) |
| Ethnicity framework | ONS 19 categories (race-based) | Visible minority + ethnic origin (self-reported, multi-response) |
| Religion data | Census 2021 (voluntary question) | Census 2021 (25% sample, mandatory long-form) |
| Language dimension | Minimal (Welsh, Gaelic, other) | Critical: English/French bilingualism + 200+ mother tongues |
| Indigenous dimension | Not applicable | Essential: First Nations, Metis, Inuit (5% of population) |
| Health system | NHS (single system) | Medicare (13 provincial/territorial systems, same principles) |
| Media landscape | BBC-centred, national tabloids | CBC/Radio-Canada, regional chains, French/English split |
| Political parties | Labour/Conservative/LibDem/SNP/etc. | Liberal/Conservative/NDP/BQ/Green (BQ Quebec-only) |
| Spending data | ONS Family Spending | StatCan SHS (similar provincial granularity) |
| Housing data | ONS + Land Registry | Census + CMHC + CREA |
