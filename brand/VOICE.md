# Voice (Sprachstil)

How trackIT Systems writes, in German and English. Applies to the website, product texts, datasheets, newsletters and emails.

Status: **draft**. The decisions under "Settled" are fixed. "How we write" describes the voice of the current website and newsletter and still needs to be confirmed in review.

## Settled

### Address / Anrede

- Default register is neutral and academic: we describe systems, methods and results, mostly without addressing the reader.
- When we address the reader directly, German uses **Sie**. English uses "you".
- Record: [ADR-0004](docs/adr/0004-address-and-gender.md).

### Gender-inclusive language / Gendern

- German: neutral wording where it is common ("Forschende", "Teilnehmende", "Fachleute"). Where no common neutral form exists, use the gender star ("Kund\*innen").
- English: neutral by default ("they", "researchers").
- Record: [ADR-0004](docs/adr/0004-address-and-gender.md).

### Languages

- German and English are both first-class. Each text is written natively in its language, not translated word by word.
- English is American English (spelling, vocabulary, punctuation). Record: [ADR-0008](docs/adr/0008-american-english.md).

## How we write (derived from trackit.systems and the newsletter)

Derived from the website and six newsletter issues (DE and EN, 2026) on 2026-11-02. This section describes the voice we already have and want to keep. Quotes are from the site; the full texts are in [examples/voice/](examples/voice/). To be confirmed in the review session.

In one sentence: researchers who build the tools they needed in the field, explaining them to colleagues. Precise about method and numbers, open about limits, warm and plain when we talk about ourselves.

### Explain the method before the product

Product pages follow the same path: what the method is, how it is done today, what the problem is, what our system changes. The radio telemetry page has the sections „Die Radiotelemetrie“, „Manuelle Radiotelemetrie“, „Was genau ist das Problem?“, „Automatische Radiotelemetrie“.

> Für die Erfassung von Raumbewegungen kleiner Tiere werden daher seit den 1960er Jahren Radiosender verwendet, die in regelmäßigen Abständen ein Signal aussenden. Anhand der Signalstärke kann mit Richtantennen die Richtung bestimmt werden, aus der das Signal kommt.

Examples: [good-de-manuelle-radiotelemetrie](examples/voice/good-de-manuelle-radiotelemetrie.md).

### Use the vocabulary of the field

We write for ecologists, consultants and authorities, and use their terms without simplifying: Kreuzpeilung, Senderlaufzeit, Quartiernutzung, Prädation, flügge, Limikolen, Multilateration, Yagi-Antennen, Klangattrappe. Scientific names go in titles („Ziegenmelker (Caprimulgus europaeus)“). Abbreviations are spelled out once („VHF-Sendern (very high frequency)“). Method descriptions go as deep as a reviewer needs. Product and tool names follow the [glossary](glossary/terms.yml): devices are `<domain>:tracker` with variants after a dash (audio:tracker-monitor, vhf:tracker-directional), the system for one domain is audio:pipeline or vhf:pipeline, everything together is „das trackIT System“, the operating system is tsOS, online tools have their own names (EcoHub, Sonari).

Examples: [good-de-methode-directional](examples/voice/good-de-methode-directional.md).

### Numbers, units and ranges

Claims come with figures: how many, how long, how precise, where, since when, at what price.

> Die Arbeitszeit reduziert sich auf die Installation der Sensorik (ca. 1 Stunde pro Einheit), ermöglicht jedoch die kontinuierliche Aufzeichnung von bis zu 60 Individuen.

> Für diese geplante Überwachung der Bestandsentwicklungen und Populationstrends kommen insgesamt 65 unserer audio:tracker zum Einsatz. Diese Systeme werden dauerhaft im Gelände installiert und verbleiben über den gesamten Projektzeitraum von 2026 bis 2031 vor Ort.

Ranges are given, not only the best value („lag die Überlebenswahrscheinlichkeit lediglich zwischen 0 und 27 %“). Product variants are lists of facts („20h Laufzeit“, „Ladung via USB-C“, „Mietpreis pro Saison: 1200€“).

Examples: [good-de-automatische-radiotelemetrie-zahlen](examples/voice/good-de-automatische-radiotelemetrie-zahlen.md), [good-de-produktdaten](examples/voice/good-de-produktdaten.md).

### Start conservation topics with the species

Project pages and posts open with the state of the species, in figures. Our motivation, biodiversity loss, shows through these facts.

> Die Bestände der Uferschnepfe in Deutschland haben sich mehr als halbiert, die Zahl der Kiebitze ist in Deutschland seit den 1980er Jahren um 93% gesunken, der Brachvogel war bis in das 20. Jahrhundert ein häufiger Vogel der Hoch- und Niedermoore und ist heute vom Aussterben bedroht.

Examples: [good-de-wiesenbrueter-einstieg](examples/voice/good-de-wiesenbrueter-einstieg.md).

### Questions guide the reader

Headlines, section headings and calls to action are often questions, followed directly by the answer: „Was genau ist das Problem?“, „Worum geht es?“, „Warum sterben trotz erfolgreichem Schlupf weiterhin so viele Küken?“, „Neugierig geworden?“, „Weitere Fragen zur Radiotelemetrie?“.

### Say what the data shows, and what it does not

Inferences are marked as such, other people's experience is attributed, and the system's role is kept in proportion.

> Da 63 % dieser Küken während der Nacht oder in der Dämmerung verschwanden, liegt nahe, dass auch hier überwiegend nächtliche Raubsäuger verantwortlich waren.

> Nach Angaben des Projektteams erleichtert die kontinuierliche Verfügbarkeit der Positionsdaten sowohl die Lokalisierung von Nestern als auch das gezielte Wiederfangen besenderter Individuen.

> Die fachliche Bewertung der Ergebnisse bleibt Aufgabe der Bearbeitenden. Das System dient als Werkzeug zur Datenerhebung und entlastet insbesondere bei zeitaufwendigen Routineaufgaben.

> Ziel des Projekts ist nicht, etablierte Kartier- und Monitoringmethoden zu ersetzen. Vielmehr sollen automatische Erfassungsverfahren bestehende Methoden ergänzen und zusätzliche Informationen liefern.

Our methods complement fieldwork and expert judgment; we do not present them as a replacement. Results that are not yet analyzed are called that („Die erhobenen Daten werden jetzt ausgewertet und analysiert.“), and first results include their exceptions („Tiere mit eingeschränkter Sehfunktion zeigen hierbei eine Störung.“).

Examples: [good-de-ziegenmelker-grenzen](examples/voice/good-de-ziegenmelker-grenzen.md), [good-de-newsletter-hlnug-ergaenzen](examples/voice/good-de-newsletter-hlnug-ergaenzen.md), [good-de-newsletter-auffangstationen](examples/voice/good-de-newsletter-auffangstationen.md), [good-de-audiotracker-beispiel](examples/voice/good-de-audiotracker-beispiel.md) (validation by the Vogelschutzwarte).

### One everyday example for an abstract feature

> So muss zum Beispiel nicht jedes erkannte Rotkehlchen oder jede erkannte Zwergfledermaus als Aufnahme-Datei gespeichert werden.

### Name partners and credit people

Institutions with their full name, project leads with title, photo credits under images, links to the partner's report or press release. Field work gets acknowledged: „Diese Arbeit verdient großen Respekt.“ Customers speak in their own words, with name and role, in testimonials.

### Figures are described like in a paper

Captions say what is plotted, what the symbols mean, and what to see, including what the manual method missed.

> An Tag eins und vier wurde eine händische Quartiersuche durchgeführt (Diamantsymbole zeigen die tatsächliche Distanz des gefundenen Quartiers). Der Quartierwechsel an Tag drei in den 200 m Radius von Station vier wurde dabei nicht festgestellt.

Examples: [good-de-abbildungen](examples/voice/good-de-abbildungen.md).

### Structure of a project post

1. Lead: the species or the problem, in two or three sentences.
2. The question the project answers.
3. The project: who, where, when, how many animals.
4. How the system helps: a list of „Begriff: Satz.“ items („Gezielte Kontrollen im Gelände: Aktuelle Positions- und Empfangsdaten können aus der Ferne eingesehen werden.“).
5. Results, as numbers.
6. Credits and links to further material (report, press, TV piece).
7. The standard newsletter line.

Newsletter issues add an „Ausblick“ that turns results into practical steps for conservation work („Rückzugsräume: Ungemähte Vegetationsinseln als Versteck anbieten.“).

Example: [good-de-brachvogel-beitrag](examples/voice/good-de-brachvogel-beitrag.md), [good-de-newsletter-ausblick-praxis](examples/voice/good-de-newsletter-ausblick-praxis.md).

### Newsletter

The newsletter („Projekt des Monats“) uses the same voice as the website, addressed directly to the reader with Sie.

- **Greeting:** „Liebe Partner\*innen und Interessent\*innen,“.
- **Opening:** one short paragraph on what is in this issue, in a modest tone („ein spannendes Projekt …, das wir mit unserer Technologie unterstützen dürfen“). Or start straight with the research question as a subheading.
- **Subheadings as questions** lead through the topic: „Wie funktioniert das genau?“, „Warum ist das interessant?“, „Was bedeutet das für Sie?“.
- **One main project** per issue, built like a project post (see above), with the partner quoted by name and title and photo credits with the names of the people shown.
- **Short items** after it: move, media coverage, dates („Nächste Termine“), each with the facts first.
- **„Ausblick“** on the next issue, one or two sentences.
- **Closing:** an invitation to share one's own project („Ihr Projekt in unserem Newsletter?“), then the standard lines: „Sie kennen jemanden, der sich für moderne Naturschutz-Technologien oder unsere Projekte interessiert? Gerne darf dieser Newsletter weitergeleitet werden.“

Examples: [good-de-newsletter-rahmen](examples/voice/good-de-newsletter-rahmen.md), [good-de-newsletter-hlnug-ergaenzen](examples/voice/good-de-newsletter-hlnug-ergaenzen.md), [good-de-newsletter-auffangstationen](examples/voice/good-de-newsletter-auffangstationen.md).

### About ourselves: plain, warm, a little humor

News about the company, workshops and events are personal and light, and still give all dates, times and addresses. Exclamation marks in news headlines are fine („Gewonnen!“, „Wir sind dabei!“). Motivation is told in a person's own words, not claimed for the company.

> … und freuen uns darauf, den Höhlenstatus hinter uns zu lassen und endlich wieder Tageslicht durch echte Fenster zu sehen!

> Der Workshop startet um 09:00 Uhr und wird, je nachdem wie schnell wir durchkommen, bis in den späten Nachmittag/frühen Abend hinein dauern.

> “Wir stecken mitten im sechsten Massenaussterben und um Arten effektiv zu schützen, müssen wir erst einmal fundiert verstehen, was da draußen überhaupt passiert.”

Team bios give the path and the tasks, with the first name the person goes by („Chris“, „Tobi“, „Melli“).

Examples: [good-de-ueber-uns-ton](examples/voice/good-de-ueber-uns-ton.md), [good-de-team-bios](examples/voice/good-de-team-bios.md).

### „trackIT Systems“ and „wir“

Descriptive pages speak of „trackIT Systems“ (the company) and „das trackIT System“ (everything we provide) in the third person; where a text means only radio telemetry or only acoustics, it names the vhf:pipeline or audio:pipeline; posts, news and calls to action use „wir“. Every page ends with a named contact person and their role.

### English

The newer English posts (curlew, nightjar) are the model: same structure and tone as the German original, written as English, with the field's terms (tagged, fledged, birds of prey).

Example: [good-en-curlew-results](examples/voice/good-en-curlew-results.md).

## Where website and newsletter fall short of their own voice

A few passages do not match the rest. They are the exceptions, not the tone to change.

- **Generic words where the page has the specifics.** One paragraph about „Lösungen“ is copied onto four pages; „innovativ“, „effizient“, „Komplettpaket“ appear where the concrete version is a few lines away. „Lösung(en)“ appears 26 times on the German site, „solution(s)“ 23 times in English. See [bad-de-loesungen-boilerplate](examples/voice/bad-de-loesungen-boilerplate.md).
- **Sales block after a careful text.** The newsletter section „Was bedeutet das für Sie?“ follows an honest account of the method with „Skalierbare Lösungen für Projekte jeder Größe“ and two more generic promises. See [bad-de-newsletter-was-bedeutet-das](examples/voice/bad-de-newsletter-was-bedeutet-das.md).
- **Comparatives without the comparison.** „deutlich besser und präziser“ in the summary of the Eingriffsverfahren page, while the comparison study sits in Abb. 3 on the same page. See [bad-de-eingriffsverfahren-steigerungen](examples/voice/bad-de-eingriffsverfahren-steigerungen.md).
- **Stacked praise in the award post.** The joy fits; „Engagement … nachhaltiger und innovativer Lösungen“, „wertvolle Bestätigung“, „neue Chancen“ do not. See [bad-de-gruenderpreis-selbstlob](examples/voice/bad-de-gruenderpreis-selbstlob.md).
- **Machine-translated English on the older product pages.** „Recording of train path queries“ for Trassenquerungen, „neighborhood use“ for Quartiernutzung, long German noun chains. The English newsletters are better but still translated: „Dear Partners and Interested Parties“, „emergency departments“ for Auffangstationen, mixed British and American spelling (we use American English, see ADR-0008). See [bad-en-radiotelemetry-translation](examples/voice/bad-en-radiotelemetry-translation.md), [bad-en-services-one-sentence](examples/voice/bad-en-services-one-sentence.md), [bad-en-newsletter-translation](examples/voice/bad-en-newsletter-translation.md).
- **Uniqueness claims.** „weltweit das einzige“ (radio telemetry page), „derzeit das einzige verfügbare System“ (nightjar post and newsletter), „weltweit erstmalig“ and „eine einzigartige Möglichkeit“ (HLNUG newsletter) need a qualifier or a source. TODO(brand owner): decide whether we keep the claim.
- **Typos and spelling of our own name.** „Aufösung“, „Besenstil“, „Dabei diesem“, „trackIT systems“ with a lowercase s. Not a voice issue, but it costs the same credibility; every newsletter needs a second reader.
- **Address and gender forms.** The settled rules above are not followed everywhere: ihr/euch on the workshop page and the move notice (switching to Sie within one list), du on the jobs page and the Girls' Day post, and one „könnt ihr“ in an otherwise Sie newsletter (event notice in the Auffangstationen issue). Gender forms appear in six spellings: *Kund\*innen, Anwender\*Innen, Forscher\*Innen, PraktikerInnen, Gründer:innen, Kunden*. See [bad-de-anrede-ihr-du](examples/voice/bad-de-anrede-ihr-du.md). TODO(brand owner): decide whether job ads and offers for school students may use du; that extends ADR-0004 and needs a new ADR.

Words that show up where the specifics are missing, for a future Vale check: DE *Lösung(en), innovativ, effizient, optimal, Komplettpaket*; EN *solution(s), innovative*. Words that do not appear on the site and should stay out: DE *nahtlos, maßgeschneidert, ganzheitlich, revolutionär, Gamechanger, eintauchen, entfesseln, Mehrwert schaffen*; EN *seamless, delve, leverage, cutting-edge, unlock, empower, elevate, game-changer*.

## How we derive the rest

1. Collect real texts in [examples/voice/](examples/voice/): our own (good and bad), and from organizations whose tone we want to be close to.
2. For each, note what works or fails, in one or two lines.
3. Derive rules from the patterns; each rule gets a before/after pair from a real text.
4. Settle disputed rules with an ADR.

TODO(brand owner): schedule the review session; decide who contributes examples.
