package data

// Nouns is an array of nouns
var Nouns = []string{"Armour", "Barrymore", "Cabot", "Catholicism", "Chihuahua", "Christianity", "Easter", "Frenchman", "Lowry", "Mayer", "Orientalism", "Pharaoh", "Pueblo", "Pullman", "Rodeo", "Saturday", "Sister", "Snead", "Syrah", "Tuesday", "Woodward", "abbey", "absence", "absorption", "abstinence", "absurdity", "abundance", "acceptance", "accessibility", "accommodation", "accomplice", "accountability", "accounting", "accreditation", "accuracy", "acquiescence", "acreage", "actress", "actuality", "adage", "adaptation", "adherence", "adjustment", "adoption", "adultery", "advancement", "advert", "advertisement", "advertising", "advice", "aesthetics", "affinity", "aggression", "agriculture", "aircraft", "airtime", "allegation", "allegiance", "allegory", "allergy", "allies", "alligator", "allocation", "allotment", "altercation", "ambulance", "ammonia", "anatomy", "anemia", "ankle", "announcement", "annoyance", "annuity", "anomaly", "anthropology", "anxiety", "apartheid", "apologise", "apostle", "apparatus", "appeasement", "appellation", "appendix", "applause", "appointment", "appraisal", "archery", "archipelago", "architecture", "ardor", "arrears", "arrow", "artisan", "artistry", "ascent", "assembly", "assignment", "association", "asthma", "atheism", "attacker", "attraction", "attractiveness", "auspices", "authority", "avarice", "aversion", "aviation", "babbling", "backlash", "baker", "ballet", "balls", "banjo", "baron", "barrier", "barrister", "bases", "basin", "basis", "battery", "battling", "bedtime", "beginner", "begun", "bending", "bicycle", "billing", "bingo", "biography", "biology", "birthplace", "blackberry", "blather", "blossom", "boardroom", "boasting", "bodyguard", "boldness", "bomber", "bondage", "bonding", "bones", "bonus", "bookmark", "boomer", "booty", "bounds", "bowling", "brainstorming", "breadth", "breaker", "brewer", "brightness", "broccoli", "broth", "brotherhood", "browsing", "brunch", "brunt", "building", "bullion", "bureaucracy", "burglary", "buyout", "by-election", "cabal", "cabbage", "calamity", "campaign", "canonization", "captaincy", "carcass", "carrier", "cartridge", "cassette", "catfish", "caught", "celebrity", "cemetery", "certainty", "certification", "charade", "chasm", "check-in", "cheerleader", "cheesecake", "chemotherapy", "chili", "china", "chivalry", "cholera", "cilantro", "circus", "civilisation", "civility", "clearance", "clearing", "clerk", "climber", "closeness", "clothing", "clutches", "coaster", "coconut", "coding", "collaborator", "colleague", "college", "collision", "colors", "combustion", "comedian", "comer", "commander", "commemoration", "commenter", "commissioner", "commune", "competition", "completeness", "complexity", "computing", "comrade", "concur", "condominium", "conduit", "confidant", "configuration", "confiscation", "conflagration", "conflict", "consist", "consistency", "consolidation", "conspiracy", "constable", "consul", "consultancy", "contentment", "contents", "contractor", "conversation", "cornerstone", "corpus", "correlation", "councilman", "counselor", "countdown", "countryman", "coverage", "covering", "coyote", "cracker", "creator", "criminality", "crocodile", "cropping", "cross-examination", "crossover", "crossroads", "culprit", "cumin", "curator", "curfew", "cursor", "custard", "cutter", "cyclist", "cyclone", "cylinder", "cynicism", "daddy", "damsel", "darkness", "dawning", "daybreak", "dealing", "dedication", "deduction", "defection", "deference", "deficiency", "definition", "deflation", "degeneration", "delegation", "delicacy", "delirium", "deliverance", "demeanor", "demon", "demonstration", "denomination", "dentist", "departure", "depletion", "depression", "designation", "despotism", "detention", "developer", "devolution", "dexterity", "diagnosis", "dialect", "differentiation", "digger", "digress", "dioxide", "diploma", "disability", "disarmament", "discord", "discovery", "dishonesty", "dismissal", "disobedience", "dispatcher", "disservice", "distribution", "distributor", "diver", "diversity", "docking", "dollar", "dominance", "domination", "dominion", "donkey", "doorstep", "doorway", "dossier", "downside", "drafting", "drank", "drilling", "driver", "drumming", "drunkenness", "duchess", "ducking", "dugout", "dumps", "dwelling", "dynamics", "eagerness", "earnestness", "earnings", "eater", "editor", "effectiveness", "electricity", "elements", "eloquence", "emancipation", "embodiment", "embroidery", "emperor", "employment", "encampment", "enclosure", "encouragement", "endangerment", "enlightenment", "enthusiasm", "environment", "environs", "envoy", "epilepsy", "equation", "equator", "error", "espionage", "estimation", "evacuation", "exaggeration", "examination", "exclamation", "expediency", "exploitation", "extinction", "eyewitness", "falls", "fascism", "fastball", "feces", "feedback", "ferocity", "fertilization", "fetish", "finale", "firing", "fixing", "flashing", "flask", "flora", "fluke", "folklore", "follower", "foothold", "footing", "forefinger", "forefront", "forgiveness", "formality", "formation", "formula", "foyer", "fragmentation", "framework", "fraud", "freestyle", "frequency", "friendliness", "fries", "frigate", "fulfillment", "function", "functionality", "fundraiser", "fusion", "futility", "gallantry", "gallery", "genesis", "genitals", "girlfriend", "glamour", "glitter", "glucose", "google", "grandeur", "grappling", "greens", "gridlock", "grocer", "groundwork", "grouping", "gunman", "gusto", "habitation", "hacker", "hallway", "hamburger", "hammock", "handling", "hands", "handshake", "happiness", "hardship", "headcount", "header", "headquarters", "heads", "headset", "hearth", "hearts", "heath", "hegemony", "height", "hello", "helper", "helping", "helplessness", "hierarchy", "hoarding", "hockey", "homeland", "homer", "honesty", "horror", "horseman", "hostility", "housing", "humility", "hurricane", "iceberg", "ignition", "illness", "illustration", "illustrator", "immunity", "immunization", "imperialism", "imprisonment", "inaccuracy", "inaction", "inactivity", "inauguration", "indecency", "indicator", "inevitability", "infamy", "infiltration", "influx", "iniquity", "innocence", "innovation", "insanity", "inspiration", "instruction", "instructor", "insurer", "interact", "intercession", "intercourse", "intermission", "interpretation", "intersection", "interval", "intolerance", "intruder", "invasion", "investment", "involvement", "irrigation", "iteration", "jenny", "jogging", "jones", "joseph", "juggernaut", "juncture", "jurisprudence", "juror", "kangaroo", "kingdom", "knocking", "laborer", "larceny", "laurels", "layout", "leadership", "leasing", "legislation", "leopard", "liberation", "licence", "lifeblood", "lifeline", "ligament", "lighting", "likeness", "line-up", "lineage", "liner", "lineup", "liquidation", "listener", "literature", "litigation", "litre", "loathing", "locality", "lodging", "logic", "longevity", "lookout", "lordship", "lustre", "ma'am", "machinery", "madness", "magnificence", "mahogany", "mailing", "mainframe", "maintenance", "majority", "manga", "mango", "manifesto", "mantra", "manufacturer", "maple", "martin", "martyrdom", "mathematician", "matrix", "matron", "mayhem", "mayor", "means", "meantime", "measurement", "mechanics", "mediator", "medics", "melodrama", "memory", "mentality", "metaphysics", "method", "metre", "miner", "mirth", "misconception", "misery", "mishap", "misunderstanding", "mobility", "molasses", "momentum", "monarchy", "monument", "morale", "mortality", "motto", "mouthful", "mouthpiece", "mover", "movie", "mowing", "murderer", "musician", "mutation", "mythology", "narration", "narrator", "nationality", "negligence", "neighborhood", "neighbour", "nervousness", "networking", "nexus", "nightmare", "nobility", "nobody", "noodle", "normalcy", "notification", "nourishment", "novella", "nucleus", "nuisance", "nursery", "nutrition", "nylon", "oasis", "obscenity", "obscurity", "observer", "offense", "onslaught", "operation", "opportunity", "opposition", "oracle", "orchestra", "organisation", "organizer", "orientation", "originality", "ounce", "outage", "outcome", "outdoors", "outfield", "outing", "outpost", "outset", "overseer", "owner", "oxygen", "pairing", "panther", "paradox", "parliament", "parsley", "parson", "passenger", "pasta", "patchwork", "pathos", "patriotism", "pendulum", "penguin", "permission", "persona", "perusal", "pessimism", "peter", "philosopher", "phosphorus", "phrasing", "physique", "piles", "plateau", "playing", "plaza", "plethora", "plurality", "pneumonia", "pointer", "poker", "policeman", "polling", "poster", "posterity", "posting", "postponement", "potassium", "pottery", "poultry", "pounding", "pragmatism", "precedence", "precinct", "preoccupation", "pretense", "priesthood", "prisoner", "privacy", "probation", "proceeding", "proceedings", "processing", "processor", "progression", "projection", "prominence", "propensity", "prophecy", "prorogation", "prospectus", "protein", "prototype", "providence", "provider", "provocation", "proximity", "puberty", "publicist", "publicity", "publisher", "pundit", "putting", "quantity", "quart", "quilting", "quorum", "racism", "radiance", "ralph", "rancher", "ranger", "rapidity", "rapport", "ratification", "rationality", "reaction", "reader", "reassurance", "rebirth", "receptor", "recipe", "recognition", "recourse", "recreation", "rector", "recurrence", "redemption", "redistribution", "redundancy", "refinery", "reformer", "refrigerator", "regularity", "regulator", "reinforcement", "reins", "reinstatement", "relativism", "relaxation", "rendition", "repayment", "repentance", "repertoire", "repository", "republic", "reputation", "resentment", "residency", "resignation", "restaurant", "resurgence", "retailer", "retention", "retirement", "reviewer", "riches", "righteousness", "roadblock", "robber", "rocks", "rubbing", "runoff", "saloon", "salvation", "sarcasm", "saucer", "savior", "scarcity", "scenario", "scenery", "schism", "scholarship", "schoolboy", "schooner", "scissors", "scolding", "scooter", "scouring", "scrimmage", "scrum", "seating", "sediment", "seduction", "seeder", "seizure", "self-confidence", "self-control", "self-respect", "semicolon", "semiconductor", "semifinal", "senator", "sending", "serenity", "seriousness", "servitude", "sesame", "setup", "sewing", "sharpness", "shaving", "shoplifting", "shopping", "siding", "simplicity", "simulation", "sinking", "skate", "sloth", "slugger", "snack", "snail", "snapshot", "snark", "soccer", "solemnity", "solicitation", "solitude", "somewhere", "sophistication", "sorcery", "souvenir", "spaghetti", "specification", "specimen", "specs", "spectacle", "spectre", "speculation", "sperm", "spoiler", "squad", "squid", "staging", "stagnation", "staircase", "stairway", "stamina", "standpoint", "standstill", "stanza", "statement", "stillness", "stimulus", "stocks", "stole", "stoppage", "storey", "storyteller", "stylus", "subcommittee", "subscription", "subsidy", "suburb", "success", "sufferer", "supposition", "suspension", "sweater", "sweepstakes", "swimmer", "syndrome", "synopsis", "syntax", "system", "tablespoon", "taker", "tavern", "technology", "telephony", "template", "tempo", "tendency", "tendon", "terrier", "terror", "terry", "theater", "theology", "therapy", "thicket", "thoroughfare", "threshold", "thriller", "thunderstorm", "ticker", "tiger", "tights", "to-day", "tossing", "touchdown", "tourist", "tourney", "toxicity", "tracing", "tractor", "translation", "transmission", "transmitter", "trauma", "traveler", "treadmill", "trilogy", "trout", "tuning", "twenties", "tycoon", "tyrant", "ultimatum", "underdog", "underwear", "unhappiness", "unification", "university", "uprising", "vaccination", "validity", "vampire", "vanguard", "variation", "vegetation", "verification", "viability", "vicinity", "victory", "viewpoint", "villa", "vindication", "violation", "vista", "vocalist", "vogue", "volcano", "voltage", "vomiting", "vulnerability", "waistcoat", "waitress", "wardrobe", "warmth", "watchdog", "wealth", "weariness", "whereabouts", "whisky", "whiteness", "widget", "width", "windfall", "wiring", "witchcraft", "withholding", "womanhood", "words", "workman", "youngster"}
