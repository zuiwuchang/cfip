{
  // 間隔多久執行一次任務，如果不設置則只執行一次
  // interval: '6h',

  // 並行工作數量，建議不要設置太高，太高的並行存在諸多問題
  // 1. 對測試的網站造成太大壓力，並且可能會被加入黑名單禁止訪問
  // 2. 容易引發防火牆的 sni 阻斷，一段時間內都禁止你和 cloudflare 的任何連接
  // 3. 對運行測試的計算機造成過大的壓力
  // 默認 5 就很不錯，無論對運行 測試的計算機 cloudflare 測試網站造成的影響都很微弱，也不容易引發 sni 阻斷
  worker: 5,
  request: {
    // 請求超時時間
    timeout: '400ms',

    // 對同一 ip 要進行測試多少次
    count: 3,

    // 請求 url
    url: 'https://usa.visa.com/',

    // 只要請求返回的 http status 與此相等時才認爲請求成功，
    // 如果 code <1 則 將 [200,299] 的 htttp status 都認爲成功
    // code:200,
    // 防止 tls 阻斷
    sni: [
      'www.visa.com.hk',
      'kw.visamiddleeast.com',  // ar-KW 阿拉伯語(科威特)
      'www.visa.co.in',  // en-IN 英語(印度)
      'www.visa.com.bo',  // es-BO 西班牙語(玻利維亞)
      'ae.visamiddleeast.com',  // en-AE
      'www.visa.co.za',  // en-ZA 英語(南非)
      'www.visa.no',  // no-NO
      'www.visa.gr',  // el-GR
      'www.visa.nl',  // nl-NL
      'www.visa.com.au',  // en-AU
      'www.visa.co.il',  // he-IL

      'https://www.visa.hu',  // hu-HU
      'www.visa.com.jm',  // en-JM
      'www.visa.com.sv',  // es-SV
      'bd.visa.com',  // en-BD
      'www.visa.co.cr',  // es-CR
      'www.visa.cl',  // es-CL
      'www.visa.ca',  // fr-CA
      'www.visa.com.co',  // es-CO
      'caribbean.visa.com',  // en-BL
      'www.visa.pl',  // pl-PL
      'www.visa.pt',  // pt-PT
      'sa.visamiddleeast.com',  // en-SA
      'eg.visamiddleeast.com',  // ar-EG
      'www.visa.be',  // fr-BE
      'www.visa.com.sg',  // en-SG
      'ma.visamiddleeast.com',  // ar-MA
      'cis.visa.com',  // ru-TJ
      'www.visa.com.bz',  // en-BZ

      'www.visa.co.ke',  // en-KE
      'www.visa.com.hn',  // es-HN
      'www.visa.com.ge',  // en-GE
      'www.visa.com.hr',  // hr-HR
      'www.visa.com.kz',  // ru-KZ
      'www.visa.com.pr',  // es-PR
      'www.visa.com.py',  // es-PY
      'www.visa.gp',  // en-GP
      'www.visaeurope.at',  // de-AT
      'www.visa.is',  // is-IS
      'www.visabg.com',  // bg-BG
      'www.visa.cz',  // cs-CZ
      'www.visa.com.ph',  // en-PH
      'by.visa.com',  // ru-BY
      'pk.visamiddleeast.com',  // en-PK
      'www.visa.com.tw"',  // zh-TW
      'www.visakorea.com',  // ko-KR

      'www.visa.sk',  // sk-SK
      'www.visa.com.ar',  // es-AR
      'qa.visamiddleeast.com',  // en-QA
      'www.visa.se',  // sv-SE
      'www.visa.dk',  // da-DK
      'www.visa.mn',  // mn-MN
      'www.visa.com.az',  // ru-AZ
      'www.visa.com.ua',  // uk-UA
      'usa.visa.com',  // en-US
      'myanmar.visa.com',  // en-MM
      'www.visa.lv',  // lv-LV
      'www.visa.mq',  // en-MQ
      'www.visa.com.mt',  // en-MT
      // 'www.visa.cn',  // zh-CN
      'africa.visa.com',  // en-MW

      'www.visa.com.my',  // en-MY
      'www.visa.com.ng',  // en-NG
      'www.visa.com.vn',  // en-VN
      'www.visa.com.cy',  // el-CY
      'www.visa.co.jp',  // ja-JP
      'www.visaeurope.ch',  // de-CH
      'www.visa.com.gt',  // es-GT
      'www.visa.com.pe',  // es-PE
      'www.visa.co.nz',  // en-NZ
      'www.visa.com.pa',  // es-PA
      'www.visa.ro',  // ro-RO
      'www.visa.com.tr',  // tr-TR
      'www.visa.ee',  // et-EE
      'www.visa.com.kh',  // km-KM
      'www.visa.fr',  // fr-FR

      'www.visa.co.uk',  //  en-GB
      'www.visa.fi',  // fi-FI
      'km.visamiddleeast.com',  // en-KM
      'www.visa.lt',  // lt-LT
      'www.visaeurope.si',  // sl-SI
      'www.visa.com.do',  // es-DO
      'kw.visamiddleeast.com',  // en-KW
      'www.visa.co.th',  // en-TH
      'www.visa.ie',  // ie-GB
      'www.visa.com.lk',  // en-LK
      'www.visa.de',  //  de-DE
      'www.visa.com.tt',  // en-TT
      'rs.visa.com',  // sr-RS
      'www.visa.com.br',  // pt-BR
      'www.visa.com.uy',  // es-UY
      'www.visa.es',  // es-ES
      'www.visa.co.ve',  // es-VE
      'www.visaeurope.lu',  // en-LU
      'www.visa.com.mx',  // es-MX
      'www.visaitalia.com',  // it-IT
      'www.visa.co.id',  // id-ID

      'www.visasoutheasteurope.com',  // en-ME
      // es-NI
      'www.visa.co.ni',
    ],
    // 如果設置每次都以不同的 UserAgent 發送請求
    userAgent: [
      'Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/113.0',
      'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36',

      'Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:131.0) Gecko/20100101 Firefox/131.0',
      'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36',
    ],
  },
  found: {
    // 每次任務的目標是要尋找多少個 ip
    ip: 10,
    // 至少要尋找多少個有效 ip，從中選擇最優(平均延遲最低) ip
    valid: 50,
    // 至少要測試多少個 ip
    test: 1000,
  },
}
