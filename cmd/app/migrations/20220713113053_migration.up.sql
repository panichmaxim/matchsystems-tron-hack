create table category
(
    id             bigserial
        primary key,
    name           varchar not null,
    risk           bigint  not null,
    description_ru varchar not null,
    description_en varchar not null
);

INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (1, 'Darknet Marketplace', 100, '', 'A website with high-risk activities');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (2, 'Darknet Service', 100, '', 'An individual participant in a high-risk activity');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (3, 'Drugs', 100, '', 'An individual participant in high-risk drugs activities');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (4, 'Precursors', 85, '', 'An individual participant in high-risk activities specifically in the area of sale of precursors (goods and substances used in the creation of drugs)');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (5, 'China precursors manufacturing', 85, '', 'Precursor manufacturer from China');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (6, 'Child Abuse', 100, '', 'Distribution of prohibited materials involving minors');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (7, 'Stolen Credit Cards', 100, '', 'Funds received by carders as a result of illegal activities');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (8, 'Blackmail', 100, '', 'Ransom letters under the guise of hacking a computer or spreading defamatory information');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (9, 'Personal data RU', 100, '', 'Sale of personal data of citizens of the Russian Federation');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (10, 'Personal data CIS', 100, '', 'Sale of personal data of citizens of other CIS countries (except the Russian Federation)');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (11, 'Personal data EU', 100, '', 'Sale of personal data of EU citizens');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (12, 'Personal data US', 100, '', 'Sale of personal data of US citizens');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (13, 'Personal data other', 100, '', 'Sale of personal data of citizens of other states');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (14, 'Stolen DataBase', 100, '', 'Sale of databases (not penetrations, but the databases themselves)');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (15, 'Smuggling drugs', 100, '', 'Participant in drug smuggling');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (16, 'Smuggling precursors', 100, '', 'Participant in the smuggling of precursors');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (17, 'Smuggling people', 100, '', 'Participant in people smuggling');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (18, 'Smuggling weapons', 100, '', 'weapons smuggler');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (19, 'Smuggling', 100, '', 'Participant in the smuggling of other goods');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (20, 'Illegal migration', 100, '', 'Organizer of illegal migration');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (21, 'Human trafficking', 100, '', 'Human trafficking (without smuggling)');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (22, 'Fake documents', 100, '', 'Making fake documents (not renderings, but physical documents)');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (23, 'Fake document rendering', 100, '', 'Manufacture of forged documents (not renderings, but physical documents)');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (24, 'Illegal weapons', 100, '', 'Sale of weapons');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (25, 'Laundering of money', 100, '', 'Laundering of proceeds from crime');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (26, 'Illegal financial transactions', 100, '', 'Illegal financial transactions');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (27, 'Exchange With High ML Risk', 75, '', 'Cryptocurrency exchange with a high risk of money laundering');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (28, 'Exchange With Low ML Risk', 25, '', 'Cryptocurrency exchange with a low risk of money laundering');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (29, 'Exchange With Moderate ML Risk', 50, '', 'Cryptocurrency exchange with an average risk of money laundering');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (30, 'Exchange With Very High ML Risk', 100, '', 'Cryptocurrency exchange with a very high risk of money laundering');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (31, 'Fraudulent Exchange', 100, '', 'Fake cryptocurrency exchange (a kind of scam)');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (32, 'Gambling', 100, '', 'Gambling activity');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (33, 'Illegal Service', 100, '', 'Other forms of illegal activity not specified in other categories');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (34, 'Miner', 0, '', 'Miner');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (35, 'Mixing Service', 100, '', 'Cryptocurrency mixer');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (36, 'Online Marketplace', 50, '', 'Marketplace');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (37, 'Online Wallet', 50, '', 'Online cryptocurrency wallet');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (38, 'PEP', 75, '', 'Politically exposed person');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (39, 'Corruption', 100, '', 'Corruption');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (40, 'Bridge', 50, '', 'Cryptocurrency bridge (like renBTC, WBTC, etc.)');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (41, 'DEX (excluding Bridges)', 50, '', 'Deposit addresses of various DEXs');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (42, 'P2P Exchange With High ML Risk', 75, '', 'P2P exchange operator with high risk of money laundering');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (43, 'P2P Exchange With Low ML Risk', 25, '', 'P2P exchange operator with low risk of money laundering');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (44, 'P2P Exchange With Moderate ML Risk', 50, '', 'P2P exchange operator with medium risk of money laundering');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (45, 'P2P Exchange With Very High ML Risk', 90, '', 'P2P exchange operator with a very high risk of money laundering');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (46, 'Payment Processor', 50, '', 'Cryptocurrency payment processor');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (47, 'Ransom', 100, '', 'Malware ransom');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (48, 'Malware (excluding Ransom)', 100, '', 'Developers and sellers of malware (hidden miners, installers, etc.)');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (49, 'DDOS service', 100, '', 'Providing DDOS services');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (50, 'Phishing service', 100, '', 'Phishing service');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (51, 'Scam', 100, '', 'Recipients of funds as a result of fraud (including pyramid schemes)');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (52, 'Stolen Coins', 100, '', 'Stolen cryptocurrency assets');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (53, 'Scam ICO', 100, '', 'Fake ICO');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (54, 'OTC Exchange With High ML Risk', 75, '', 'Cryptocurrency exchange with a high risk of money laundering');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (55, 'OTC Exchange With Low ML Risk', 25, '', 'Cryptocurrency exchange with low risk of money laundering');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (56, 'OTC Exchange With Moderate ML Risk', 50, '', 'Cryptocurrency exchange with medium risk of money laundering');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (57, 'OTC Exchange With Very High ML Risk', 90, '', 'Cryptocurrency exchange with a very high risk of money laundering');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (58, 'Hight Risk country', 90, '', 'FATF gray list and services of these countries');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (59, 'Sanction country', 100, '', 'Iran, North Korea, Crimea, Venezuela, Cuba and services from these countries or territories');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (60, 'Terrorism', 100, '', 'Participants and sponsors of terrorist activities');
INSERT INTO public.category (id, name, risk, description_ru, description_en) VALUES (61, 'ATM', 50, '', 'Cryptocurrency ATM');
