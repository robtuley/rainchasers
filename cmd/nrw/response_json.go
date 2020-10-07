package main

// curl -v -X GET "https://api.naturalresources.wales/rivers-and-seas/v1/api/StationData"
//      -H "Ocp-Apim-Subscription-Key: {subscription key}"

const jsonResponseFromAPI = `
[
   {
      "coordinates":{
         "latitude":53.115798,
         "longitude":-3.2966817
      },
      "location":1000,
      "nameEN":"Brynhyfryd Ruthin School",
      "nameCY":"Brynhyfryd Ruthin School",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1000",
      "ngr":"SJ1330958410",
      "titleEn":"Brynhyfryd Ruthin School raingauge",
      "titleCy":"Mesurydd glaw Ysgol Brynhyfryd, Ruthin",
      "parameters":[
         {
            "parameter":10095,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.777033,
         "longitude":-3.3678892
      },
      "location":1001,
      "nameEN":"Clawddnewydd",
      "nameCY":"Clawddnewydd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1001",
      "ngr":"SJ0782520816",
      "titleEn":"Clawddnewydd raingauge",
      "titleCy":"Mesurydd glaw Clawddnewydd",
      "parameters":[
         {
            "parameter":10096,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.13914,
         "longitude":-3.6776504
      },
      "location":1002,
      "nameEN":"Gwytherin",
      "nameCY":"Gwytherin",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1002",
      "ngr":"SH8787361536",
      "titleEn":"Gwytherin raingauge",
      "titleCy":"Mesurydd glaw Gwytherin",
      "parameters":[
         {
            "parameter":10097,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.21974,
         "longitude":-3.2090409
      },
      "location":1003,
      "nameEN":"Moel Y Crio",
      "nameCY":"Moel Y Crio",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1003",
      "ngr":"SJ1937069870",
      "titleEn":"Moel Y Crio raingauge",
      "titleCy":"Mesurydd glaw Moel Y Crio",
      "parameters":[
         {
            "parameter":10098,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.600,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.168868,
         "longitude":-3.5703166
      },
      "location":1004,
      "nameEN":"Plas Pigot",
      "nameCY":"Plas Pigot",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1004",
      "ngr":"SH9512564680",
      "titleEn":"Plas Pigot raingauge",
      "titleCy":"Mesurydd glaw Plas Pigot",
      "parameters":[
         {
            "parameter":10099,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.264584,
         "longitude":-3.451024
      },
      "location":1005,
      "nameEN":"St Asaph",
      "nameCY":"St Asaph",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1005",
      "ngr":"SJ0331575159",
      "titleEn":"St Asaph raingauge",
      "titleCy":"Mesurydd glaw Llanelwy",
      "parameters":[
         {
            "parameter":10100,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":1.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.139911,
         "longitude":-3.8020403
      },
      "location":1006,
      "nameEN":"Llanrwst",
      "nameCY":"Llanrwst",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1006",
      "ngr":"SH7955561824",
      "titleEn":"Llanrwst raingauge",
      "titleCy":"Mesurydd glaw Llanrwst",
      "parameters":[
         {
            "parameter":10101,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.994328,
         "longitude":-3.7735611
      },
      "location":1007,
      "nameEN":"Ysbyty Ifan",
      "nameCY":"Ysbyty Ifan",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1007",
      "ngr":"SH8105945583",
      "titleEn":"Ysbyty Ifan raingauge",
      "titleCy":"Mesurydd glaw Ysbyty Ifan",
      "parameters":[
         {
            "parameter":10102,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.025858,
         "longitude":-2.9535249
      },
      "location":1008,
      "nameEN":"Five Fords",
      "nameCY":"Five Fords",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1008",
      "ngr":"SJ3614348044",
      "titleEn":"Five Fords raingauge",
      "titleCy":"Mesurydd glaw Pump Rhyd",
      "parameters":[
         {
            "parameter":10103,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.103689,
         "longitude":-3.0905588
      },
      "location":1009,
      "nameEN":"Llanfynydd",
      "nameCY":"Llanfynydd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1009",
      "ngr":"SJ2708456833",
      "titleEn":"Llanfynydd raingauge",
      "titleCy":"Mesurydd glaw Llanfynydd",
      "parameters":[
         {
            "parameter":10104,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T18:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.149188,
         "longitude":-3.1968299
      },
      "location":1010,
      "nameEN":"Loggerheads",
      "nameCY":"Loggerheads",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1010",
      "ngr":"SJ2005462008",
      "titleEn":"Loggerheads raingauge",
      "titleCy":"Mesurydd glaw Loggerheads",
      "parameters":[
         {
            "parameter":10105,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.400,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.172387,
         "longitude":-3.1374823
      },
      "location":1011,
      "nameEN":"Mold",
      "nameCY":"Mold",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1011",
      "ngr":"SJ2406464524",
      "titleEn":"Mold raingauge",
      "titleCy":"Mesurydd glaw Yr Wyddgrug",
      "parameters":[
         {
            "parameter":10106,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.07142,
         "longitude":-3.1305951
      },
      "location":1012,
      "nameEN":"Nant Y Ffrith",
      "nameCY":"Nant Y Ffrith",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1012",
      "ngr":"SJ2434753285",
      "titleEn":"Nant Y Ffrith raingauge",
      "titleCy":"Mesurydd glaw Nant Y Ffrith",
      "parameters":[
         {
            "parameter":10107,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T18:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.20943,
         "longitude":-3.1254213
      },
      "location":1013,
      "nameEN":"Northop",
      "nameCY":"Northop",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1013",
      "ngr":"SJ2493568632",
      "titleEn":"Northop raingauge",
      "titleCy":"Mesurydd glaw Llaneurgain",
      "parameters":[
         {
            "parameter":10108,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.800,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.907416,
         "longitude":-3.5834893
      },
      "location":1014,
      "nameEN":"Bala Climate",
      "nameCY":"Bala Climate",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1014",
      "ngr":"SH9360235617",
      "titleEn":"Bala Climate raingauge",
      "titleCy":"Mesurydd glaw Bala Climate",
      "parameters":[
         {
            "parameter":10109,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.963273,
         "longitude":-3.7248509
      },
      "location":1015,
      "nameEN":"Cynefail",
      "nameCY":"Cynefail",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1015",
      "ngr":"SH8424542049",
      "titleEn":"Cynefail raingauge",
      "titleCy":"Mesurydd glaw Cynefail",
      "parameters":[
         {
            "parameter":10110,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.063499,
         "longitude":-3.5599289
      },
      "location":1016,
      "nameEN":"Llyn Alwen",
      "nameCY":"Llyn Alwen",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1016",
      "ngr":"SH9556452944",
      "titleEn":"Llyn Alwen raingauge",
      "titleCy":"Mesurydd glaw Llyn Alwen",
      "parameters":[
         {
            "parameter":10111,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.828567,
         "longitude":-3.7070105
      },
      "location":1017,
      "nameEN":"Pant Gwyn",
      "nameCY":"Pant Gwyn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1017",
      "ngr":"SH8508727037",
      "titleEn":"Pant Gwyn raingauge",
      "titleCy":"Mesurydd glaw Pant Gwyn",
      "parameters":[
         {
            "parameter":10112,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.945151,
         "longitude":-3.6667397
      },
      "location":1018,
      "nameEN":"Tryweryn Dam Pluvio and",
      "nameCY":"Tryweryn Dam Pluvio and",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1018",
      "ngr":"SH8810139941",
      "titleEn":"Tryweryn Dam Pluvio and raingauge",
      "titleCy":"Mesurydd glaw Argae Tryweryn",
      "parameters":[
         {
            "parameter":10113,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.812674,
         "longitude":-4.5096962
      },
      "location":1019,
      "nameEN":"Abersoch",
      "nameCY":"Abersoch",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1019",
      "ngr":"SH3095226856",
      "titleEn":"Abersoch raingauge",
      "titleCy":"Mesurydd glaw Abersoch",
      "parameters":[
         {
            "parameter":10114,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.212105,
         "longitude":-4.1423227
      },
      "location":1020,
      "nameEN":"Afon Adda",
      "nameCY":"Afon Adda",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1020",
      "ngr":"SH5703670481",
      "titleEn":"Afon Adda raingauge",
      "titleCy":"Mesurydd glaw Afon Adda",
      "parameters":[
         {
            "parameter":10115,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.0267,
         "longitude":-4.1256637
      },
      "location":1021,
      "nameEN":"Afon Colwyn",
      "nameCY":"Afon Colwyn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1021",
      "ngr":"SH5753649826",
      "titleEn":"Afon Colwyn raingauge",
      "titleCy":"Mesurydd glaw Afon Colwyn",
      "parameters":[
         {
            "parameter":10116,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.168505,
         "longitude":-4.0657308
      },
      "location":1022,
      "nameEN":"Bethesda Quarry",
      "nameCY":"Bethesda Quarry",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1022",
      "ngr":"SH6201065481",
      "titleEn":"Bethesda Quarry raingauge",
      "titleCy":"Mesurydd glaw Cloddfa Bethesda",
      "parameters":[
         {
            "parameter":10117,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.918085,
         "longitude":-4.2493954
      },
      "location":1023,
      "nameEN":"Criccieth",
      "nameCY":"Criccieth",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1023",
      "ngr":"SH4886037999",
      "titleEn":"Criccieth raingauge",
      "titleCy":"Mesurydd glaw Criccieth",
      "parameters":[
         {
            "parameter":10118,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.067313,
         "longitude":-4.0113282
      },
      "location":1024,
      "nameEN":"Cwm Dyli",
      "nameCY":"Cwm Dyli",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1024",
      "ngr":"SH6533054122",
      "titleEn":"Cwm Dyli raingauge",
      "titleCy":"Mesurydd glaw Cwm Dyli",
      "parameters":[
         {
            "parameter":10119,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.954232,
         "longitude":-4.4423451
      },
      "location":1025,
      "nameEN":"Llithfaen",
      "nameCY":"Llithfaen",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1025",
      "ngr":"SH3602642443",
      "titleEn":"Llithfaen raingauge",
      "titleCy":"Mesurydd glaw Llithfaen",
      "parameters":[
         {
            "parameter":10120,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.339472,
         "longitude":-4.4409252
      },
      "location":1026,
      "nameEN":"Llyn Alaw",
      "nameCY":"Llyn Alaw",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1026",
      "ngr":"SH3758385286",
      "titleEn":"Llyn Alaw raingauge",
      "titleCy":"Mesurydd glaw Llyn Alaw",
      "parameters":[
         {
            "parameter":10121,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T18:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.268208,
         "longitude":-4.3332764
      },
      "location":1027,
      "nameEN":"Llyn Cefni",
      "nameCY":"Llyn Cefni",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1027",
      "ngr":"SH4449077120",
      "titleEn":"Llyn Cefni raingauge",
      "titleCy":"Mesurydd glaw Llyn Cefni",
      "parameters":[
         {
            "parameter":10122,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.264115,
         "longitude":-4.6131789
      },
      "location":1028,
      "nameEN":"Trearddur Bay",
      "nameCY":"Trearddur Bay",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1028",
      "ngr":"SH2581077311",
      "titleEn":"Trearddur Bay raingauge",
      "titleCy":"Mesurydd glaw Bae Trearddur",
      "parameters":[
         {
            "parameter":10123,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.679265,
         "longitude":-3.8125868
      },
      "location":1029,
      "nameEN":"Corris - Aberllefenni",
      "nameCY":"Corris - Aberllefenni",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1029",
      "ngr":"SH7755610605",
      "titleEn":"Corris - Aberllefenni raingauge",
      "titleCy":"Mesurydd glaw Corris - Aberllefenni",
      "parameters":[
         {
            "parameter":10124,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.740274,
         "longitude":-3.6463927
      },
      "location":1030,
      "nameEN":"Llan Y Mawddwy",
      "nameCY":"Llan Y Mawddwy",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1030",
      "ngr":"SH8894617121",
      "titleEn":"Llan Y Mawddwy raingauge",
      "titleCy":"Mesurydd glaw Llan Y Mawddwy",
      "parameters":[
         {
            "parameter":10125,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.648513,
         "longitude":-3.5993058
      },
      "location":1031,
      "nameEN":"Llanbrynmair",
      "nameCY":"Llanbrynmair",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1031",
      "ngr":"SH9189806843",
      "titleEn":"Llanbrynmair raingauge",
      "titleCy":"Mesurydd glaw Llanbrynmair",
      "parameters":[
         {
            "parameter":10126,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.600,
            "latestTime":"2020-10-06T21:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.930315,
         "longitude":-4.1223199
      },
      "location":1032,
      "nameEN":"Porthmadog",
      "nameCY":"Porthmadog",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1032",
      "ngr":"SH5744339099",
      "titleEn":"Porthmadog raingauge",
      "titleCy":"Mesurydd glaw Porthmadog",
      "parameters":[
         {
            "parameter":10127,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.797615,
         "longitude":-3.785006
      },
      "location":1033,
      "nameEN":"Rhyd Y Main",
      "nameCY":"Rhyd Y Main",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1033",
      "ngr":"SH7974723722",
      "titleEn":"Rhyd Y Main raingauge",
      "titleCy":"Mesurydd glaw Rhyd Y Main",
      "parameters":[
         {
            "parameter":10128,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.600,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.809656,
         "longitude":-3.2167269
      },
      "location":1034,
      "nameEN":"Carno Reservoir",
      "nameCY":"Carno Reservoir",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1034",
      "ngr":"SO1621313036",
      "titleEn":"Carno Reservoir raingauge",
      "titleCy":"Mesurydd glaw Llyn Carno",
      "parameters":[
         {
            "parameter":10129,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.755288,
         "longitude":-3.1299887
      },
      "location":1035,
      "nameEN":"Cwmtillery",
      "nameCY":"Cwmtillery",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1035",
      "ngr":"SO2209906893",
      "titleEn":"Cwmtillery raingauge",
      "titleCy":"Mesurydd glaw Cwmtillery",
      "parameters":[
         {
            "parameter":10130,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.589311,
         "longitude":-3.1640454
      },
      "location":1036,
      "nameEN":"Machen Waterloo",
      "nameCY":"Machen Waterloo",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1036",
      "ngr":"ST1945488471",
      "titleEn":"Machen Waterloo raingauge",
      "titleCy":"Mesurydd glaw Machen Waterloo",
      "parameters":[
         {
            "parameter":10131,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.651951,
         "longitude":-3.2967281
      },
      "location":1037,
      "nameEN":"Nelson",
      "nameCY":"Nelson",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1037",
      "ngr":"ST1038695592",
      "titleEn":"Nelson raingauge",
      "titleCy":"Mesurydd glaw Ffos y Gerddinen",
      "parameters":[
         {
            "parameter":10132,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T18:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.534456,
         "longitude":-3.2272079
      },
      "location":1038,
      "nameEN":"Rhiwbina Res",
      "nameCY":"Rhiwbina Res",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1038",
      "ngr":"ST1497682442",
      "titleEn":"Rhiwbina Res raingauge",
      "titleCy":"Mesurydd glaw Llyn Rhiwbina",
      "parameters":[
         {
            "parameter":10133,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.253424,
         "longitude":-3.6387904
      },
      "location":1039,
      "nameEN":"Rhymney Reservoir",
      "nameCY":"Rhymney Reservoir",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1039",
      "ngr":"SO1048009910",
      "titleEn":"Rhymney Reservoir raingauge",
      "titleCy":"Mesurydd glaw Llyn Rhymney",
      "parameters":[
         {
            "parameter":10134,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.718341,
         "longitude":-3.0303118
      },
      "location":1040,
      "nameEN":"Trevethin",
      "nameCY":"Trevethin",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1040",
      "ngr":"SO2892102682",
      "titleEn":"Trevethin raingauge",
      "titleCy":"Mesurydd glaw Trevethin",
      "parameters":[
         {
            "parameter":10135,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.689408,
         "longitude":-3.2075244
      },
      "location":1041,
      "nameEN":"Ty Fry",
      "nameCY":"Ty Fry",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1041",
      "ngr":"ST1662699652",
      "titleEn":"Ty Fry raingauge",
      "titleCy":"Mesurydd glaw Ty Fry",
      "parameters":[
         {
            "parameter":10136,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.793003,
         "longitude":-3.4337095
      },
      "location":1042,
      "nameEN":"Llwynon Reservoir",
      "nameCY":"Llwynon Reservoir",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1042",
      "ngr":"SO0121811456",
      "titleEn":"Llwynon Reservoir raingauge",
      "titleCy":"Mesurydd glaw Llwynon Reservoir",
      "parameters":[
         {
            "parameter":10137,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.679039,
         "longitude":-3.4881214
      },
      "location":1043,
      "nameEN":"Maerdy WTW",
      "nameCY":"Maerdy WTW",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1043",
      "ngr":"SS9720798857",
      "titleEn":"Maerdy WTW raingauge",
      "titleCy":"Mesurydd glaw Maerdy WTW",
      "parameters":[
         {
            "parameter":10138,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.657853,
         "longitude":-3.3975177
      },
      "location":1044,
      "nameEN":"Nant Yr Ysfa",
      "nameCY":"Nant Yr Ysfa",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1044",
      "ngr":"ST0342696377",
      "titleEn":"Nant Yr Ysfa raingauge",
      "titleCy":"Mesurydd glaw Nant Yr Ysfa",
      "parameters":[
         {
            "parameter":10139,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.797889,
         "longitude":-3.3662566
      },
      "location":1045,
      "nameEN":"Pontsticill Upper",
      "nameCY":"Pontsticill Upper",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1045",
      "ngr":"SO0588011910",
      "titleEn":"Pontsticill Upper raingauge",
      "titleCy":"Mesurydd glaw Pontsticill Uchaf",
      "parameters":[
         {
            "parameter":10140,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.48559,
         "longitude":-3.2672733
      },
      "location":1046,
      "nameEN":"St Fagans",
      "nameCY":"St Fagans",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1046",
      "ngr":"ST1210377055",
      "titleEn":"St Fagans raingauge",
      "titleCy":"Mesurydd glaw Sain Ffagan",
      "parameters":[
         {
            "parameter":10141,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.684092,
         "longitude":-3.2987023
      },
      "location":1047,
      "nameEN":"Trelewis",
      "nameCY":"Trelewis",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1047",
      "ngr":"ST1031399169",
      "titleEn":"Trelewis raingauge",
      "titleCy":"Mesurydd glaw Trelewis",
      "parameters":[
         {
            "parameter":10142,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.681388,
         "longitude":-3.5444086
      },
      "location":1048,
      "nameEN":"Tyn Y Waun Treherbert",
      "nameCY":"Tyn Y Waun Treherbert",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1048",
      "ngr":"SS9332199199",
      "titleEn":"Tyn Y Waun Treherbert raingauge",
      "titleCy":"Mesurydd glaw Tyn Y Waun Treherbert",
      "parameters":[
         {
            "parameter":10143,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.939018,
         "longitude":-3.4015996
      },
      "location":1049,
      "nameEN":"Brecon",
      "nameCY":"Brecon",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1049",
      "ngr":"SO0374527652",
      "titleEn":"Brecon raingauge",
      "titleCy":"Mesurydd glaw Aberhonddu",
      "parameters":[
         {
            "parameter":10144,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.577237,
         "longitude":-2.8022824
      },
      "location":1050,
      "nameEN":"Collister Climate Station",
      "nameCY":"Collister Climate Station",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1050",
      "ngr":"ST4450186791",
      "titleEn":"Collister Climate Station raingauge",
      "titleCy":"Mesurydd glaw Collister Climate Station",
      "parameters":[
         {
            "parameter":10145,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.887655,
         "longitude":-3.6229202
      },
      "location":1051,
      "nameEN":"Crai Reservoir",
      "nameCY":"Crai Reservoir",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1051",
      "ngr":"SN8840422256",
      "titleEn":"Crai Reservoir raingauge",
      "titleCy":"Mesurydd glaw Crai Reservoir",
      "parameters":[
         {
            "parameter":10146,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T21:30:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.757223,
         "longitude":-3.060275
      },
      "location":1052,
      "nameEN":"Cwmafon",
      "nameCY":"Cwmafon",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1052",
      "ngr":"SO2691407036",
      "titleEn":"Cwmafon raingauge",
      "titleCy":"Mesurydd glaw Cwmafon",
      "parameters":[
         {
            "parameter":10147,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.615371,
         "longitude":-2.9808499
      },
      "location":1053,
      "nameEN":"Lodge Farm Frog",
      "nameCY":"Lodge Farm Frog",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1053",
      "ngr":"ST3218491183",
      "titleEn":"Lodge Farm Frog raingauge",
      "titleCy":"Mesurydd glaw Lodge Farm Frog",
      "parameters":[
         {
            "parameter":10148,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.400,
            "latestTime":"2020-10-06T16:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.647988,
         "longitude":-3.0549345
      },
      "location":1054,
      "nameEN":"Maes Y Rhiw",
      "nameCY":"Maes Y Rhiw",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1054",
      "ngr":"ST2710794882",
      "titleEn":"Maes Y Rhiw raingauge",
      "titleCy":"Mesurydd glaw Maes Y Rhiw",
      "parameters":[
         {
            "parameter":10149,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.872726,
         "longitude":-3.4792559
      },
      "location":1055,
      "nameEN":"Storey Arms",
      "nameCY":"Storey Arms",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1055",
      "ngr":"SN9825720385",
      "titleEn":"Storey Arms raingauge",
      "titleCy":"Mesurydd glaw Storey Arms",
      "parameters":[
         {
            "parameter":10150,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.926712,
         "longitude":-3.1276293
      },
      "location":1056,
      "nameEN":"Tal Y Maes",
      "nameCY":"Tal Y Maes",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1056",
      "ngr":"SO2255725956",
      "titleEn":"Tal Y Maes raingauge",
      "titleCy":"Mesurydd glaw Tal Y Maes",
      "parameters":[
         {
            "parameter":10151,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.067549,
         "longitude":-3.4469381
      },
      "location":1057,
      "nameEN":"Upper Chapel No1",
      "nameCY":"Upper Chapel No1",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1057",
      "ngr":"SO0091342008",
      "titleEn":"Upper Chapel No1 raingauge",
      "titleCy":"Mesurydd glaw Capel Uchaf",
      "parameters":[
         {
            "parameter":10152,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.946841,
         "longitude":-3.69791
      },
      "location":1058,
      "nameEN":"Usk Reservoir No2",
      "nameCY":"Usk Reservoir No2",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1058",
      "ngr":"SN8339728956",
      "titleEn":"Usk Reservoir No2 raingauge",
      "titleCy":"Mesurydd glaw Llyn Wysg",
      "parameters":[
         {
            "parameter":10153,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.600,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.594504,
         "longitude":-3.0319675
      },
      "location":1059,
      "nameEN":"Ynysfro",
      "nameCY":"Ynysfro",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1059",
      "ngr":"ST2861288911",
      "titleEn":"Ynysfro raingauge",
      "titleCy":"Mesurydd glaw Ynysfro",
      "parameters":[
         {
            "parameter":10154,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.419039,
         "longitude":-3.2095882
      },
      "location":1060,
      "nameEN":"Cogmoors",
      "nameCY":"Cogmoors",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1060",
      "ngr":"ST1598669586",
      "titleEn":"Cogmoors raingauge",
      "titleCy":"Mesurydd glaw Cogmoors",
      "parameters":[
         {
            "parameter":10155,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.453264,
         "longitude":-3.4452457
      },
      "location":1061,
      "nameEN":"Cowbridge STW",
      "nameCY":"Cowbridge STW",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1061",
      "ngr":"SS9967573689",
      "titleEn":"Cowbridge STW raingauge",
      "titleCy":"Mesurydd glaw Bont Faen",
      "parameters":[
         {
            "parameter":10156,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.554746,
         "longitude":-3.4032689
      },
      "location":1062,
      "nameEN":"Dyffryn Isaf, Llantrisant",
      "nameCY":"Dyffryn Isaf, Llantrisant",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1062",
      "ngr":"ST0280884918",
      "titleEn":"Dyffryn Isaf, Llantrisant raingauge",
      "titleCy":"Mesurydd glaw Dyffryn Isaf, Llantrisant",
      "parameters":[
         {
            "parameter":10157,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.534929,
         "longitude":-3.3377777
      },
      "location":1063,
      "nameEN":"Rhiwsaeson STW",
      "nameCY":"Rhiwsaeson STW",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1063",
      "ngr":"ST0730882629",
      "titleEn":"Rhiwsaeson STW raingauge",
      "titleCy":"Mesurydd glaw Rhiwsaeson STW",
      "parameters":[
         {
            "parameter":10158,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.101146,
         "longitude":-3.61863
      },
      "location":1064,
      "nameEN":"Abernant",
      "nameCY":"Abernant",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1064",
      "ngr":"SN8922845993",
      "titleEn":"Abernant raingauge",
      "titleCy":"Mesurydd glaw Abernant",
      "parameters":[
         {
            "parameter":10159,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.800,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.301424,
         "longitude":-3.1696106
      },
      "location":1065,
      "nameEN":"Bleddfa",
      "nameCY":"Bleddfa",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1065",
      "ngr":"SO2034267678",
      "titleEn":"Bleddfa raingauge",
      "titleCy":"Mesurydd glaw Bleddfa",
      "parameters":[
         {
            "parameter":10160,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.232077,
         "longitude":-3.2010548
      },
      "location":1066,
      "nameEN":"Break Your Neck Falls",
      "nameCY":"Break Your Neck Falls",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1066",
      "ngr":"SO1807060000",
      "titleEn":"Break Your Neck Falls raingauge",
      "titleCy":"Mesurydd glaw Break Your Neck Falls",
      "parameters":[
         {
            "parameter":10161,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.361103,
         "longitude":-3.4251721
      },
      "location":1067,
      "nameEN":"Bwlch y Sarnau",
      "nameCY":"Bwlch y Sarnau",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1067",
      "ngr":"SO0304774628",
      "titleEn":"Bwlch y Sarnau Raingauge",
      "titleCy":"Mesurydd glaw Bwlch y Sarnau",
      "parameters":[
         {
            "parameter":10162,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.253548,
         "longitude":-3.6389414
      },
      "location":1068,
      "nameEN":"Ciloerwynt",
      "nameCY":"Ciloerwynt",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1068",
      "ngr":"SN8822062974",
      "titleEn":"Ciloerwynt raingauge",
      "titleCy":"Mesurydd glaw Ciloerwynt",
      "parameters":[
         {
            "parameter":10163,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.274175,
         "longitude":-3.563221
      },
      "location":1069,
      "nameEN":"Elan Village",
      "nameCY":"Elan Village",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1069",
      "ngr":"SN9343865154",
      "titleEn":"Elan Village raingauge",
      "titleCy":"Mesurydd glaw Elan Village",
      "parameters":[
         {
            "parameter":10164,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.390532,
         "longitude":-3.3284839
      },
      "location":1070,
      "nameEN":"Llanbadarn Fynydd",
      "nameCY":"Llanbadarn Fynydd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1070",
      "ngr":"SO0969177776",
      "titleEn":"Llanbadarn Fynydd raingauge",
      "titleCy":"Mesurydd glaw Llanbadarn Fynydd",
      "parameters":[
         {
            "parameter":10165,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.83357,
         "longitude":-3.6817437
      },
      "location":1071,
      "nameEN":"Llanerch Yrfa",
      "nameCY":"Llanerch Yrfa",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1071",
      "ngr":"SN8361955503",
      "titleEn":"Llanerch Yrfa raingauge",
      "titleCy":"Mesurydd glaw Llanerch Yrfa",
      "parameters":[
         {
            "parameter":10166,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.404212,
         "longitude":-3.6068589
      },
      "location":1072,
      "nameEN":"Llangurig",
      "nameCY":"Llangurig",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1072",
      "ngr":"SN9078279682",
      "titleEn":"Llangurig raingauge",
      "titleCy":"Mesurydd glaw Llangurig",
      "parameters":[
         {
            "parameter":10167,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":1.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.251331,
         "longitude":-3.3906523
      },
      "location":1073,
      "nameEN":"Llanyre",
      "nameCY":"Llanyre",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1073",
      "ngr":"SO0516362373",
      "titleEn":"Llanyre raingauge",
      "titleCy":"Mesurydd glaw Llanyre",
      "parameters":[
         {
            "parameter":10168,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.428262,
         "longitude":-3.7013381
      },
      "location":1074,
      "nameEN":"Station 1074 Name EN",
      "nameCY":"Station 1074 Name CY",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1074",
      "ngr":"SN8441882504",
      "titleEn":"Station 1074 HoverTitle EN",
      "titleCy":"Station 1074 HoverTitle CY",
      "parameters":[
         {
            "parameter":10169,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.800,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.295116,
         "longitude":-3.4987181
      },
      "location":1075,
      "nameEN":"Rhyader STW",
      "nameCY":"Rhyader STW",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1075",
      "ngr":"SN9788767390",
      "titleEn":"Rhyader STW raingauge",
      "titleCy":"Mesurydd glaw Rhyader STW",
      "parameters":[
         {
            "parameter":10170,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.959281,
         "longitude":-3.053554
      },
      "location":1076,
      "nameEN":"Station 1076 Name EN",
      "nameCY":"Station 1076 Name CY",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1076",
      "ngr":"SO2770329502",
      "titleEn":"Station 1076 HoverTitle EN",
      "titleCy":"Station 1076 HoverTitle CY",
      "parameters":[
         {
            "parameter":10171,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.004936,
         "longitude":-3.236804
      },
      "location":1077,
      "nameEN":"Talgarth",
      "nameCY":"Talgarth",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1077",
      "ngr":"SO1519834778",
      "titleEn":"Talgarth raingauge",
      "titleCy":"Mesurydd glaw Talgarth",
      "parameters":[
         {
            "parameter":10172,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.033022,
         "longitude":-3.1738457
      },
      "location":1078,
      "nameEN":"Tregoyd",
      "nameCY":"Tregoyd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1078",
      "ngr":"SO1957037830",
      "titleEn":"Tregoyd raingauge",
      "titleCy":"Mesurydd glaw Tregoyd",
      "parameters":[
         {
            "parameter":10173,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.819674,
         "longitude":-3.8593012
      },
      "location":1079,
      "nameEN":"Brynamman",
      "nameCY":"Brynamman",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1079",
      "ngr":"SN7194415085",
      "titleEn":"Brynamman raingauge",
      "titleCy":"Mesurydd glaw Brynamman",
      "parameters":[
         {
            "parameter":10174,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:30:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.834186,
         "longitude":-3.9601068
      },
      "location":1080,
      "nameEN":"Bryncoch",
      "nameCY":"Bryncoch",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1080",
      "ngr":"SN6504016881",
      "titleEn":"Bryncoch raingauge",
      "titleCy":"Mesurydd glaw Bryncoch",
      "parameters":[
         {
            "parameter":10175,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.663436,
         "longitude":-4.1102684
      },
      "location":1081,
      "nameEN":"Bynea",
      "nameCY":"Bynea",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1081",
      "ngr":"SS5414598182",
      "titleEn":"Bynea raingauge",
      "titleCy":"Mesurydd glaw Bynea",
      "parameters":[
         {
            "parameter":10176,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.699999,
         "longitude":-4.1489598
      },
      "location":1082,
      "nameEN":"Felinfoel",
      "nameCY":"Felinfoel",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1082",
      "ngr":"SN5158902326",
      "titleEn":"Felinfoel raingauge",
      "titleCy":"Mesurydd glaw Felinfoel",
      "parameters":[
         {
            "parameter":10177,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.770371,
         "longitude":-4.0004888
      },
      "location":1083,
      "nameEN":"Garnswllt",
      "nameCY":"Garnswllt",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1083",
      "ngr":"SN6206309860",
      "titleEn":"Garnswllt raingauge",
      "titleCy":"Mesurydd glaw Garnswllt",
      "parameters":[
         {
            "parameter":10178,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.800,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.811983,
         "longitude":-4.0850684
      },
      "location":1084,
      "nameEN":"Gorslas",
      "nameCY":"Gorslas",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1084",
      "ngr":"SN5636014651",
      "titleEn":"Gorslas raingauge",
      "titleCy":"Mesurydd glaw Gorslas",
      "parameters":[
         {
            "parameter":10179,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.800749,
         "longitude":-4.2424719
      },
      "location":1085,
      "nameEN":"Station 1085 Name EN",
      "nameCY":"Station 1085 Name CY",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1085",
      "ngr":"SN4547213724",
      "titleEn":"Station 1085 HoverTitle EN",
      "titleCy":"Station 1085 HoverTitle CY",
      "parameters":[
         {
            "parameter":10180,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.735134,
         "longitude":-3.939118
      },
      "location":1086,
      "nameEN":"Upper Lliw",
      "nameCY":"Upper Lliw",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1086",
      "ngr":"SN6619305827",
      "titleEn":"Upper Lliw raingauge",
      "titleCy":"Mesurydd glaw Lliw Uchaf",
      "parameters":[
         {
            "parameter":10181,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.400,
            "latestTime":"2020-10-06T20:30:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.679626,
         "longitude":-3.6546251
      },
      "location":1087,
      "nameEN":"Abercregan",
      "nameCY":"Abercregan",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1087",
      "ngr":"SS8569799170",
      "titleEn":"Abercregan raingauge",
      "titleCy":"Mesurydd glaw Abercregan",
      "parameters":[
         {
            "parameter":10182,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.600,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.720656,
         "longitude":-3.7544117
      },
      "location":1088,
      "nameEN":"Crynant",
      "nameCY":"Crynant",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1088",
      "ngr":"SN7890803894",
      "titleEn":"Crynant raingauge",
      "titleCy":"Mesurydd glaw Creunant",
      "parameters":[
         {
            "parameter":10183,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.400,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.661741,
         "longitude":-3.7060829
      },
      "location":1089,
      "nameEN":"Fforch Dwm",
      "nameCY":"Fforch Dwm",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1089",
      "ngr":"SS8209397263",
      "titleEn":"Fforch Dwm raingauge",
      "titleCy":"Mesurydd glaw Fforch Dwm",
      "parameters":[
         {
            "parameter":10184,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":1.400,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.749142,
         "longitude":-3.5423442
      },
      "location":1090,
      "nameEN":"Hirwaun",
      "nameCY":"Hirwaun",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1090",
      "ngr":"SN9362306731",
      "titleEn":"Hirwaun raingauge",
      "titleCy":"Mesurydd glaw Hirwaun",
      "parameters":[
         {
            "parameter":10185,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.70922,
         "longitude":-3.7023257
      },
      "location":1091,
      "nameEN":"Resolven",
      "nameCY":"Resolven",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1091",
      "ngr":"SN8247602537",
      "titleEn":"Resolven raingauge",
      "titleCy":"Mesurydd glaw Resolfen",
      "parameters":[
         {
            "parameter":10186,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.635761,
         "longitude":-3.5252862
      },
      "location":1092,
      "nameEN":"Blaenogwr",
      "nameCY":"Blaenogwr",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1092",
      "ngr":"SS9453794097",
      "titleEn":"Blaenogwr raingauge",
      "titleCy":"Mesurydd glaw Blaenogwr",
      "parameters":[
         {
            "parameter":10187,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.504862,
         "longitude":-3.5398578
      },
      "location":1093,
      "nameEN":"Brackla",
      "nameCY":"Brackla",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1093",
      "ngr":"SS9322279561",
      "titleEn":"Brackla raingauge",
      "titleCy":"Mesurydd glaw Bracla",
      "parameters":[
         {
            "parameter":10188,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.644625,
         "longitude":-3.6452708
      },
      "location":1094,
      "nameEN":"Croeserw",
      "nameCY":"Croeserw",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1094",
      "ngr":"SS8625695263",
      "titleEn":"Croeserw raingauge",
      "titleCy":"Mesurydd glaw Croeserw",
      "parameters":[
         {
            "parameter":10189,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":1.200,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.581417,
         "longitude":-3.6196189
      },
      "location":1095,
      "nameEN":"Llety Brongu",
      "nameCY":"Llety Brongu",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1095",
      "ngr":"SS8787588194",
      "titleEn":"Llety Brongu raingauge",
      "titleCy":"Mesurydd glaw Llety Brongu",
      "parameters":[
         {
            "parameter":10190,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":1.600,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.760836,
         "longitude":-5.0170933
      },
      "location":1096,
      "nameEN":"Bolton Hill",
      "nameCY":"Bolton Hill",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1096",
      "ngr":"SM9188611214",
      "titleEn":"Bolton Hill raingauge",
      "titleCy":"Mesurydd glaw Bolton Hill",
      "parameters":[
         {
            "parameter":10191,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.851752,
         "longitude":-4.8887233
      },
      "location":1097,
      "nameEN":"Clarbeston Road",
      "nameCY":"Clarbeston Road",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1097",
      "ngr":"SN0114420964",
      "titleEn":"Clarbeston Road raingauge",
      "titleCy":"Mesurydd glaw Treglarbes",
      "parameters":[
         {
            "parameter":10192,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.949257,
         "longitude":-4.591997
      },
      "location":1098,
      "nameEN":"Llanfyrnach",
      "nameCY":"Llanfyrnach",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1098",
      "ngr":"SN2196331037",
      "titleEn":"Llanfyrnach raingauge",
      "titleCy":"Mesurydd glaw Llanfyrnach",
      "parameters":[
         {
            "parameter":10193,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.933469,
         "longitude":-4.812237
      },
      "location":1099,
      "nameEN":"Maenclochog",
      "nameCY":"Maenclochog",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1099",
      "ngr":"SN0676229844",
      "titleEn":"Maenclochog raingauge",
      "titleCy":"Mesurydd glaw Maenclochog",
      "parameters":[
         {
            "parameter":10194,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.942333,
         "longitude":-5.0822954
      },
      "location":1100,
      "nameEN":"Mathry",
      "nameCY":"Mathry",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1100",
      "ngr":"SM8824231582",
      "titleEn":"Mathry raingauge",
      "titleCy":"Mesurydd glaw Mathry",
      "parameters":[
         {
            "parameter":10195,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.720185,
         "longitude":-4.7250113
      },
      "location":1101,
      "nameEN":"Pentlepoir",
      "nameCY":"Pentlepoir",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1101",
      "ngr":"SN1187005900",
      "titleEn":"Pentlepoir raingauge",
      "titleCy":"Mesurydd glaw Pentlepoir",
      "parameters":[
         {
            "parameter":10196,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.786916,
         "longitude":-4.6363572
      },
      "location":1102,
      "nameEN":"Tavernspite",
      "nameCY":"Tavernspite",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1102",
      "ngr":"SN1826113095",
      "titleEn":"Tavernspite raingauge",
      "titleCy":"Mesurydd glaw Tafarnspite",
      "parameters":[
         {
            "parameter":10197,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.67497,
         "longitude":-4.7279362
      },
      "location":1103,
      "nameEN":"Tenby",
      "nameCY":"Tenby",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1103",
      "ngr":"SN1148000880",
      "titleEn":"Tenby raingauge",
      "titleCy":"Mesurydd glaw Dinbych-y-Pysgod",
      "parameters":[
         {
            "parameter":10198,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.944699,
         "longitude":-4.5026199
      },
      "location":1104,
      "nameEN":"Trelech",
      "nameCY":"Trelech",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1104",
      "ngr":"SN2808730315",
      "titleEn":"Trelech raingauge",
      "titleCy":"Mesurydd glaw Trelech",
      "parameters":[
         {
            "parameter":10199,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.456274,
         "longitude":-3.939291
      },
      "location":1105,
      "nameEN":"Bontgoch",
      "nameCY":"Bontgoch",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1105",
      "ngr":"SN6832486027",
      "titleEn":"Bontgoch raingauge",
      "titleCy":"Mesurydd glaw Bontgoch",
      "parameters":[
         {
            "parameter":10200,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.438418,
         "longitude":-4.032387
      },
      "location":1106,
      "nameEN":"Bow Street",
      "nameCY":"Bow Street",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1106",
      "ngr":"SN6194384215",
      "titleEn":"Bow Street raingauge",
      "titleCy":"Mesurydd glaw Bow Street",
      "parameters":[
         {
            "parameter":10201,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.396077,
         "longitude":-3.899963
      },
      "location":1107,
      "nameEN":"Station 1107 Name EN",
      "nameCY":"Station 1107 Name CY",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1107",
      "ngr":"SN7082079261",
      "titleEn":"Station 1107 HoverTitle EN",
      "titleCy":"Station 1107 HoverTitle CY",
      "parameters":[
         {
            "parameter":10202,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.43132,
         "longitude":-3.8445516
      },
      "location":1108,
      "nameEN":"Dinas",
      "nameCY":"Dinas",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1108",
      "ngr":"SN7469083083",
      "titleEn":"Dinas raingauge",
      "titleCy":"Mesurydd glaw Dinas",
      "parameters":[
         {
            "parameter":10203,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":1.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.361068,
         "longitude":-3.8778697
      },
      "location":1109,
      "nameEN":"Frongoch",
      "nameCY":"Frongoch",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1109",
      "ngr":"SN7222275328",
      "titleEn":"Frongoch raingauge",
      "titleCy":"Mesurydd glaw Frongoch",
      "parameters":[
         {
            "parameter":10204,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.35856,
         "longitude":-3.8033513
      },
      "location":1110,
      "nameEN":"Pwll Peiran",
      "nameCY":"Pwll Peiran",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1110",
      "ngr":"SN7728974920",
      "titleEn":"Pwll Peiran raingauge",
      "titleCy":"Mesurydd glaw Pwll Peiran",
      "parameters":[
         {
            "parameter":10205,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.675909,
         "longitude":-3.8676233
      },
      "location":1111,
      "nameEN":"Birchgrove",
      "nameCY":"Birchgrove",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1111",
      "ngr":"SS7096199112",
      "titleEn":"Birchgrove raingauge",
      "titleCy":"Mesurydd glaw Y Gellifedw",
      "parameters":[
         {
            "parameter":10206,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.831788,
         "longitude":-3.682432
      },
      "location":1112,
      "nameEN":"Dan Yr Ogof",
      "nameCY":"Dan Yr Ogof",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1112",
      "ngr":"SN8416516136",
      "titleEn":"Dan Yr Ogof raingauge",
      "titleCy":"Mesurydd glaw Dan Yr Ogof",
      "parameters":[
         {
            "parameter":10207,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.653591,
         "longitude":-4.0321915
      },
      "location":1113,
      "nameEN":"Gowerton",
      "nameCY":"Gowerton",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1113",
      "ngr":"SS5951496934",
      "titleEn":"Gowerton raingauge",
      "titleCy":"Mesurydd glaw Tregwyr",
      "parameters":[
         {
            "parameter":10208,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.566161,
         "longitude":-4.2683361
      },
      "location":1114,
      "nameEN":"Pitton",
      "nameCY":"Pitton",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1114",
      "ngr":"SS4287887693",
      "titleEn":"Pitton raingauge",
      "titleCy":"Mesurydd glaw Pitton",
      "parameters":[
         {
            "parameter":10209,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.570137,
         "longitude":-4.0840957
      },
      "location":1115,
      "nameEN":"Southgate",
      "nameCY":"Southgate",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1115",
      "ngr":"SS5565987755",
      "titleEn":"Southgate raingauge",
      "titleCy":"Mesurydd glaw Southgate",
      "parameters":[
         {
            "parameter":10210,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.718996,
         "longitude":-3.9246881
      },
      "location":1116,
      "nameEN":"Spyte",
      "nameCY":"Spyte",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1116",
      "ngr":"SN6714204006",
      "titleEn":"Spyte raingauge",
      "titleCy":"Mesurydd glaw Spyte",
      "parameters":[
         {
            "parameter":10211,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.400,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.705465,
         "longitude":-3.8649473
      },
      "location":1117,
      "nameEN":"Trebanos",
      "nameCY":"Trebanos",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1117",
      "ngr":"SN7123002394",
      "titleEn":"Trebanos raingauge",
      "titleCy":"Mesurydd glaw Trebanos",
      "parameters":[
         {
            "parameter":10212,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":1.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.61241,
         "longitude":-3.962119
      },
      "location":1118,
      "nameEN":"Victoria Park",
      "nameCY":"Victoria Park",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1118",
      "ngr":"SS6423892222",
      "titleEn":"Victoria Park raingauge",
      "titleCy":"Mesurydd glaw Parc Victoria",
      "parameters":[
         {
            "parameter":10213,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.818137,
         "longitude":-3.5388527
      },
      "location":1119,
      "nameEN":"Ystradfellte",
      "nameCY":"Ystradfellte",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1119",
      "ngr":"SN9402614399",
      "titleEn":"Ystradfellte raingauge",
      "titleCy":"Mesurydd glaw Ystradfellte",
      "parameters":[
         {
            "parameter":10214,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.76356,
         "longitude":-3.7769427
      },
      "location":1120,
      "nameEN":"Ystradgynlais",
      "nameCY":"Ystradgynlais",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1120",
      "ngr":"SN7746808703",
      "titleEn":"Ystradgynlais raingauge",
      "titleCy":"Mesurydd glaw Ystradgynlais",
      "parameters":[
         {
            "parameter":10215,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.784139,
         "longitude":-3.8212604
      },
      "location":1121,
      "nameEN":"Ystradowen",
      "nameCY":"Ystradowen",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1121",
      "ngr":"SN7446711067",
      "titleEn":"Ystradowen raingauge",
      "titleCy":"Mesurydd glaw Ystradowen",
      "parameters":[
         {
            "parameter":10216,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.122017,
         "longitude":-4.0428775
      },
      "location":1122,
      "nameEN":"Cellan",
      "nameCY":"Cellan",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1122",
      "ngr":"SN6023749047",
      "titleEn":"Cellan raingauge",
      "titleCy":"Mesurydd glaw Cellan",
      "parameters":[
         {
            "parameter":10217,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.112845,
         "longitude":-4.6550283
      },
      "location":1123,
      "nameEN":"Ferwig",
      "nameCY":"Ferwig",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1123",
      "ngr":"SN1829749384",
      "titleEn":"Ferwig raingauge",
      "titleCy":"Mesurydd glaw Ferwig",
      "parameters":[
         {
            "parameter":10218,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.062139,
         "longitude":-4.5848751
      },
      "location":1124,
      "nameEN":"Llechryd",
      "nameCY":"Llechryd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1124",
      "ngr":"SN2289943572",
      "titleEn":"Llechryd raingauge",
      "titleCy":"Mesurydd glaw Llechryd",
      "parameters":[
         {
            "parameter":10219,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.246116,
         "longitude":-4.0669889
      },
      "location":1125,
      "nameEN":"Penuwch",
      "nameCY":"Penuwch",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1125",
      "ngr":"SN5898062895",
      "titleEn":"Penuwch raingauge",
      "titleCy":"Mesurydd glaw Penuwch",
      "parameters":[
         {
            "parameter":10220,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.284198,
         "longitude":-3.8660693
      },
      "location":1126,
      "nameEN":"Pontrydfendigaid No2",
      "nameCY":"Pontrydfendigaid No2",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1126",
      "ngr":"SN7280566758",
      "titleEn":"Pontrydfendigaid No2 raingauge",
      "titleCy":"Mesurydd glaw Pontrydfendigaid No2",
      "parameters":[
         {
            "parameter":10221,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":1.600,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.12733,
         "longitude":-4.2502918
      },
      "location":1127,
      "nameEN":"Rhos Ymryson",
      "nameCY":"Rhos Ymryson",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1127",
      "ngr":"SN4605750058",
      "titleEn":"Rhos Ymryson raingauge",
      "titleCy":"Mesurydd glaw Rhos Ymryson",
      "parameters":[
         {
            "parameter":10222,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.98461,
         "longitude":-4.0599328
      },
      "location":1128,
      "nameEN":"Abergorlech",
      "nameCY":"Abergorlech",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1128",
      "ngr":"SN5863633799",
      "titleEn":"Abergorlech raingauge",
      "titleCy":"Mesurydd glaw Abergorlech",
      "parameters":[
         {
            "parameter":10223,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.039591,
         "longitude":-3.9334614
      },
      "location":1129,
      "nameEN":"Caio",
      "nameCY":"Caio",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1129",
      "ngr":"SN6748339675",
      "titleEn":"Caio raingauge",
      "titleCy":"Mesurydd glaw Caio",
      "parameters":[
         {
            "parameter":10224,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.044935,
         "longitude":-3.7468461
      },
      "location":1130,
      "nameEN":"Cynghordy",
      "nameCY":"Cynghordy",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1130",
      "ngr":"SN8029639945",
      "titleEn":"Cynghordy raingauge",
      "titleCy":"Mesurydd glaw Cynghordy",
      "parameters":[
         {
            "parameter":10225,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.89985,
         "longitude":-3.7488421
      },
      "location":1131,
      "nameEN":"Llynyfan",
      "nameCY":"Llynyfan",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1131",
      "ngr":"SN7977123813",
      "titleEn":"Llynyfan raingauge",
      "titleCy":"Mesurydd glaw Llynyfan",
      "parameters":[
         {
            "parameter":10226,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.952713,
         "longitude":-3.7835835
      },
      "location":1132,
      "nameEN":"Myddfai",
      "nameCY":"Myddfai",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1132",
      "ngr":"SN7752529750",
      "titleEn":"Myddfai raingauge",
      "titleCy":"Mesurydd glaw Myddfai",
      "parameters":[
         {
            "parameter":10227,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.209938,
         "longitude":-3.8141568
      },
      "location":1133,
      "nameEN":"Nant Y Maen",
      "nameCY":"Nant Y Maen",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1133",
      "ngr":"SN7613958409",
      "titleEn":"Nant Y Maen raingauge",
      "titleCy":"Mesurydd glaw Nant Y Maen",
      "parameters":[
         {
            "parameter":10228,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T18:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.005407,
         "longitude":-4.2669858
      },
      "location":1134,
      "nameEN":"Pencader",
      "nameCY":"Pencader",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1134",
      "ngr":"SN4449136535",
      "titleEn":"Pencader raingauge",
      "titleCy":"Mesurydd glaw Pencader",
      "parameters":[
         {
            "parameter":10229,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.911631,
         "longitude":-4.2592189
      },
      "location":1135,
      "nameEN":"Rhydargaeau",
      "nameCY":"Rhydargaeau",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1135",
      "ngr":"SN4470026090",
      "titleEn":"Rhydargaeau raingauge",
      "titleCy":"Mesurydd glaw Rhydargaeau",
      "parameters":[
         {
            "parameter":10230,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.600,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.109203,
         "longitude":-3.7717883
      },
      "location":1136,
      "nameEN":"Ystradffin No1",
      "nameCY":"Ystradffin No1",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1136",
      "ngr":"SN7876047134",
      "titleEn":"Ystradffin No1 raingauge",
      "titleCy":"Mesurydd glaw Ystradffin No1",
      "parameters":[
         {
            "parameter":10231,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":1.600,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.414662,
         "longitude":-3.5058489
      },
      "location":1137,
      "nameEN":"Llantwit Major",
      "nameCY":"Llantwit Major",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1137",
      "ngr":"SS9537669481",
      "titleEn":"Llantwit Major raingauge",
      "titleCy":"Mesurydd glaw Llanilltud Fawr",
      "parameters":[
         {
            "parameter":10232,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.612343,
         "longitude":-3.4172196
      },
      "location":1138,
      "nameEN":"Cefn Coch",
      "nameCY":"Cefn Coch",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1138",
      "ngr":"SJ0413702562",
      "titleEn":"Cefn Coch raingauge",
      "titleCy":"Mesurydd glaw Cefn Coch",
      "parameters":[
         {
            "parameter":10233,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.500357,
         "longitude":-3.6611167
      },
      "location":1139,
      "nameEN":"Dolydd",
      "nameCY":"Dolydd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1139",
      "ngr":"SN8733790459",
      "titleEn":"Dolydd raingauge",
      "titleCy":"Mesurydd glaw Dolydd",
      "parameters":[
         {
            "parameter":10234,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.400,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.522556,
         "longitude":-3.4298627
      },
      "location":1140,
      "nameEN":"Henfryn",
      "nameCY":"Henfryn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1140",
      "ngr":"SO0308392592",
      "titleEn":"Henfryn raingauge",
      "titleCy":"Mesurydd glaw Henfryn",
      "parameters":[
         {
            "parameter":10235,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.760714,
         "longitude":-3.25586
      },
      "location":1141,
      "nameEN":"Llanfyllin",
      "nameCY":"Llanfyllin",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1141",
      "ngr":"SJ1535018863",
      "titleEn":"Llanfyllin raingauge",
      "titleCy":"Mesurydd glaw Llanfyllin",
      "parameters":[
         {
            "parameter":10236,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.82135,
         "longitude":-3.4081627
      },
      "location":1142,
      "nameEN":"Llangynog",
      "nameCY":"Llangynog",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1142",
      "ngr":"SJ0520525798",
      "titleEn":"Llangynog raingauge",
      "titleCy":"Mesurydd glaw Llangynog",
      "parameters":[
         {
            "parameter":10237,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.380125,
         "longitude":-3.5009052
      },
      "location":1143,
      "nameEN":"Nantgwyn",
      "nameCY":"Nantgwyn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1143",
      "ngr":"SN9793476848",
      "titleEn":"Nantgwyn raingauge",
      "titleCy":"Mesurydd glaw Nantgwyn",
      "parameters":[
         {
            "parameter":10238,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.400,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.717569,
         "longitude":-3.5144498
      },
      "location":1144,
      "nameEN":"Pen Y Coed",
      "nameCY":"Pen Y Coed",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1144",
      "ngr":"SH9780014400",
      "titleEn":"Pen Y Coed raingauge",
      "titleCy":"Mesurydd glaw Pen Y Coed",
      "parameters":[
         {
            "parameter":10239,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.879647,
         "longitude":-3.1591232
      },
      "location":1145,
      "nameEN":"Penygwely",
      "nameCY":"Penygwely",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1145",
      "ngr":"SJ2209131983",
      "titleEn":"Penygwely raingauge",
      "titleCy":"Mesurydd glaw Penygwely",
      "parameters":[
         {
            "parameter":10240,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Light Rain",
            "parameterStatusCY":"Glaw golau glaw",
            "units":"mm",
            "latestValue":0.200,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.320487,
         "longitude":-3.057786
      },
      "location":1146,
      "nameEN":"Rhos Y Meirch",
      "nameCY":"Rhos Y Meirch",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1146",
      "ngr":"SO2799869681",
      "titleEn":"Rhos Y Meirch raingauge",
      "titleCy":"Mesurydd glaw Rhos Y Meirch",
      "parameters":[
         {
            "parameter":10241,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.50728,
         "longitude":-3.1706195
      },
      "location":1147,
      "nameEN":"Sarn",
      "nameCY":"Sarn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1147",
      "ngr":"SO2064490576",
      "titleEn":"Sarn raingauge",
      "titleCy":"Mesurydd glaw Sarn",
      "parameters":[
         {
            "parameter":10242,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.757895,
         "longitude":-3.4604257
      },
      "location":1148,
      "nameEN":"Vyrnwy",
      "nameCY":"Vyrnwy",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1148",
      "ngr":"SJ0154018810",
      "titleEn":"Vyrnwy raingauge",
      "titleCy":"Mesurydd glaw Vyrnwy",
      "parameters":[
         {
            "parameter":10243,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.657226,
         "longitude":-3.1330043
      },
      "location":1149,
      "nameEN":"Welshpool",
      "nameCY":"Welshpool",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"1149",
      "ngr":"SJ2345907214",
      "titleEn":"Welshpool raingauge",
      "titleCy":"Mesurydd glaw Y Trallwng",
      "parameters":[
         {
            "parameter":10244,
            "paramNameEN":"Rainfall",
            "paramNameCY":"Glawiad",
            "parameterStatusEN":"Rainfall Dry",
            "parameterStatusCY":"Drws glaw",
            "units":"mm",
            "latestValue":0.000,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.759937,
         "longitude":-3.455159
      },
      "location":2003,
      "nameEN":"Vyrnwy Weir",
      "nameCY":"Cored Efyrnwy",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"2003",
      "ngr":"SJ0190019030",
      "titleEn":"Vyrnwy at Vyrnwy Weir, Llanwddyn",
      "titleCy":"Afon Efyrnwy yng Nghored Efyrnwy, Llanwddyn",
      "parameters":[
         {
            "parameter":184,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.430,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.553516,
         "longitude":-3.2335621
      },
      "location":2009,
      "nameEN":"Abermule",
      "nameCY":"Aber-miwl",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"2009",
      "ngr":"SO1646095790",
      "titleEn":"Severn at Abermule",
      "titleCy":"Afon Hafren yn Aber-miwl",
      "parameters":[
         {
            "parameter":132,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.778,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.430093,
         "longitude":-3.5457583
      },
      "location":2017,
      "nameEN":"Rhos Y Pentref",
      "nameCY":"Rhos y Pentref",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"2017",
      "ngr":"SN9500082470",
      "titleEn":"Dulas at Rhos Y Pentref, Cwmbelan",
      "titleCy":"Afon Dulas yn Rhos y Pentref, Cwmbelan",
      "parameters":[
         {
            "parameter":49,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.958,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.768986,
         "longitude":-3.108765
      },
      "location":2020,
      "nameEN":"Llanymynech",
      "nameCY":"Llanymynech",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"2020",
      "ngr":"SJ2529019620",
      "titleEn":"Vyrnwy at Llanymynech",
      "titleCy":"Afon Efyrnwy yn Llanymynech",
      "parameters":[
         {
            "parameter":185,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":2.454,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.451218,
         "longitude":-3.5391401
      },
      "location":2033,
      "nameEN":"Llanidloes",
      "nameCY":"Llanidloes",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"2033",
      "ngr":"SN9550084810",
      "titleEn":"Severn at Llanidloes",
      "titleCy":"Afon Hafren yn Llanidloes",
      "parameters":[
         {
            "parameter":133,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.546,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.514387,
         "longitude":-3.4272838
      },
      "location":2034,
      "nameEN":"Caersws",
      "nameCY":"Caersws",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"2034",
      "ngr":"SO0324091680",
      "titleEn":"Severn at Caersws",
      "titleCy":"Afon Hafren yng Nghaersws",
      "parameters":[
         {
            "parameter":134,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":2.066,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.707346,
         "longitude":-3.2501823
      },
      "location":2035,
      "nameEN":"Meifod",
      "nameCY":"Meifod",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"2035",
      "ngr":"SJ1563012920",
      "titleEn":"Vyrnwy at Meifod",
      "titleCy":"Afon Efyrnwy ym Meifod",
      "parameters":[
         {
            "parameter":186,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.608,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.468113,
         "longitude":-3.6006768
      },
      "location":2054,
      "nameEN":"Bryntail",
      "nameCY":"Bryn Tail",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"2054",
      "ngr":"SN9136086780",
      "titleEn":"Clywedog at Bryntail",
      "titleCy":"Afon Clywedog ym Mryntail",
      "parameters":[
         {
            "parameter":26,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.356,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.704591,
         "longitude":-3.3217383
      },
      "location":2060,
      "nameEN":"Pont Robert",
      "nameCY":"Pontrobert",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"2060",
      "ngr":"SJ1079012700",
      "titleEn":"Vyrnwy at Pont Robert",
      "titleCy":"Afon Efyrnwy ym Mhontrobert",
      "parameters":[
         {
            "parameter":187,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.381,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.745866,
         "longitude":-3.0410693
      },
      "location":2061,
      "nameEN":"Llandrinio",
      "nameCY":"Llandrinio",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"2061",
      "ngr":"SJ2982016980",
      "titleEn":"Severn at Llandrinio",
      "titleCy":"Afon Hafren yn Llandrinio",
      "parameters":[
         {
            "parameter":135,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":2.651,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.676678,
         "longitude":-3.4337527
      },
      "location":2062,
      "nameEN":"Llanerfyl",
      "nameCY":"Llanerfyl",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"2062",
      "ngr":"SJ0316009740",
      "titleEn":"Banwy at Llanerfyl",
      "titleCy":"Afon Banw yn Llanerfyl",
      "parameters":[
         {
            "parameter":11,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.540,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.672457,
         "longitude":-3.1157842
      },
      "location":2068,
      "nameEN":"Buttington",
      "nameCY":"Tal-y-bont",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"2068",
      "ngr":"SJ2465008890",
      "titleEn":"Severn at Buttington",
      "titleCy":"Afon Hafren yn Nhal-y-bont",
      "parameters":[
         {
            "parameter":136,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":2.538,
            "latestTime":"2020-10-06T21:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.601841,
         "longitude":-3.1668495
      },
      "location":2072,
      "nameEN":"Munlyn",
      "nameCY":"Munlyn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"2072",
      "ngr":"SJ2107001090",
      "titleEn":"Severn at Munlyn",
      "titleCy":"Afon Hafren ym Munlyn",
      "parameters":[
         {
            "parameter":137,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.986,
            "latestTime":"2020-10-06T21:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.660477,
         "longitude":-3.1556953
      },
      "location":2089,
      "nameEN":"Welshpool",
      "nameCY":"Y Trallwng",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"2089",
      "ngr":"SJ2193007600",
      "titleEn":"Lledan Brook at Welshpool",
      "titleCy":"Nant Lledan yn y Trallwng",
      "parameters":[
         {
            "parameter":88,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.208,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.101101,
         "longitude":-3.6180297
      },
      "location":4001,
      "nameEN":"Abernant",
      "nameCY":"Abernant",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4001",
      "ngr":"SN8926945987",
      "titleEn":"Irfon at Abernant",
      "titleCy":"Afon Irfon yn Abernant",
      "parameters":[
         {
            "parameter":83,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.220,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.15359,
         "longitude":-3.4117659
      },
      "location":4005,
      "nameEN":"Builth Wells",
      "nameCY":"Llanfair-ym-Muallt",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4005",
      "ngr":"SO0351051530",
      "titleEn":"Wye at Builth Wells",
      "titleCy":"Afon Gwy yn Llanfair-ym-Muallt",
      "parameters":[
         {
            "parameter":193,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":2.338,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.26906,
         "longitude":-3.5719515
      },
      "location":4008,
      "nameEN":"Caban Elan",
      "nameCY":"Caban Elan",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4008",
      "ngr":"SN9283064598",
      "titleEn":"Elan at Caban Elan",
      "titleCy":"Afon Elan yng Nghaban Elan",
      "parameters":[
         {
            "parameter":57,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.381,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.646479,
         "longitude":-2.6705531
      },
      "location":4009,
      "nameEN":"Chepstow Bridge",
      "nameCY":"Pont Cas-gwent",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4009",
      "ngr":"ST5370094400",
      "titleEn":"Wye at Chepstow Bridge",
      "titleCy":"Afon Gwy ym Mhont Cas-gwent",
      "parameters":[
         {
            "parameter":194,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":-8.674,
            "latestTime":"2020-06-09T06:30:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.1465,
         "longitude":-3.4695631
      },
      "location":4010,
      "nameEN":"Cilmery",
      "nameCY":"Cilmeri",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4010",
      "ngr":"SN9954050820",
      "titleEn":"Irfon at Cilmery",
      "titleCy":"Afon Irfon yng Nghilmeri",
      "parameters":[
         {
            "parameter":84,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.594,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.296371,
         "longitude":-3.5032918
      },
      "location":4011,
      "nameEN":"Ddol Farm",
      "nameCY":"Fferm Ddôl",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4011",
      "ngr":"SN9757867536",
      "titleEn":"Wye at Ddol Farm",
      "titleCy":"Afon Gwy yn Fferm Ddôl",
      "parameters":[
         {
            "parameter":195,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.442,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.209668,
         "longitude":-3.4289108
      },
      "location":4012,
      "nameEN":"Disserth",
      "nameCY":"Betws Diserth",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4012",
      "ngr":"SO0246057790",
      "titleEn":"Ithon at Disserth",
      "titleCy":"Afon Ieithon ym Metws Diserth",
      "parameters":[
         {
            "parameter":85,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.896,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.090631,
         "longitude":-3.3498727
      },
      "location":4013,
      "nameEN":"Erwood",
      "nameCY":"Erwyd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4013",
      "ngr":"SO0761444447",
      "titleEn":"Wye at Erwood",
      "titleCy":"Afon Gwy yn Erwyd",
      "parameters":[
         {
            "parameter":196,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":2.127,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.045649,
         "longitude":-3.1966762
      },
      "location":4014,
      "nameEN":"Glasbury",
      "nameCY":"Y Clas-ar-Wy",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4014",
      "ngr":"SO1802739260",
      "titleEn":"Wye at Glasbury",
      "titleCy":"Afon Gwy yn y Clas-ar-Wy",
      "parameters":[
         {
            "parameter":197,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.779,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.919462,
         "longitude":-2.849825
      },
      "location":4015,
      "nameEN":"Grosmont, Rhosllwyn",
      "nameCY":"Y Grysmwnt, Rhosllwyn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4015",
      "ngr":"SO4165024890",
      "titleEn":"Monnow at Grosmont, Rhosllwyn",
      "titleCy":"Afon Mynwy yn y Grysmwnt, Rhosllwyn",
      "parameters":[
         {
            "parameter":105,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.859,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.076549,
         "longitude":-3.1266888
      },
      "location":4016,
      "nameEN":"Hay On Wye",
      "nameCY":"Y Gelli Gandryll",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4016",
      "ngr":"SO2288042620",
      "titleEn":"Wye at Hay On Wye",
      "titleCy":"Afon Gwy yn y Gelli Gandryll",
      "parameters":[
         {
            "parameter":198,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.115,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.305401,
         "longitude":-3.3146545
      },
      "location":4018,
      "nameEN":"Llanddewi",
      "nameCY":"Llanddewi",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4018",
      "ngr":"SO1046068290",
      "titleEn":"Ithon at Llanddewi",
      "titleCy":"Afon Ieithon yn Llanddewi",
      "parameters":[
         {
            "parameter":86,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.557,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.151395,
         "longitude":-3.1129386
      },
      "location":4020,
      "nameEN":"Michaelchurch on Arrow",
      "nameCY":"Llanfihangel Dyffryn Arwy",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4020",
      "ngr":"SO2395050930",
      "titleEn":"Arrow at Michaelchurch on Arrow",
      "titleCy":"Afon Arwy yn Llanfihangel Dyffryn Arwy",
      "parameters":[
         {
            "parameter":10,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.479,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.797403,
         "longitude":-2.7228187
      },
      "location":4021,
      "nameEN":"Mitchel Troy",
      "nameCY":"Llanfihangel Troddi",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4021",
      "ngr":"SO5025011220",
      "titleEn":"Trothy at Mitchel Troy",
      "titleCy":"Afon Troddi yn Llanfihangel Troddi",
      "parameters":[
         {
            "parameter":163,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.419,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.30887,
         "longitude":-3.1174715
      },
      "location":4022,
      "nameEN":"Monaughty",
      "nameCY":"Monaughty",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4022",
      "ngr":"SO2391068450",
      "titleEn":"Lugg at Monaughty",
      "titleCy":"Afon Llugwy yn Monaughty",
      "parameters":[
         {
            "parameter":99,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.343,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.814631,
         "longitude":-2.703235
      },
      "location":4023,
      "nameEN":"Monmouth",
      "nameCY":"Trefynwy",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4023",
      "ngr":"SO5161913123",
      "titleEn":"Wye at Monmouth",
      "titleCy":"Afon Gwy yn Nhrefynwy",
      "parameters":[
         {
            "parameter":199,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.835,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.808838,
         "longitude":-2.7201007
      },
      "location":4024,
      "nameEN":"Monnow Gate",
      "nameCY":"Porth Trefynwy",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4024",
      "ngr":"SO5045012490",
      "titleEn":"Monnow at Monnow Gate",
      "titleCy":"Afon Mynwy ym Mhorth Trefynwy",
      "parameters":[
         {
            "parameter":106,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":2.245,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.428316,
         "longitude":-3.7013108
      },
      "location":4026,
      "nameEN":"Pant Mawr",
      "nameCY":"Pant Mawr",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4026",
      "ngr":"SN8442082510",
      "titleEn":"Wye at Pant Mawr",
      "titleCy":"Afon Gwy ym Mhant Mawr",
      "parameters":[
         {
            "parameter":200,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.540,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.796308,
         "longitude":-2.6864046
      },
      "location":4028,
      "nameEN":"Redbrook",
      "nameCY":"Redbrook",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4028",
      "ngr":"SO5276011074",
      "titleEn":"Wye at Redbrook",
      "titleCy":"Afon Gwy yn Redbrook",
      "parameters":[
         {
            "parameter":201,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.991,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.95872,
         "longitude":-3.0540211
      },
      "location":4031,
      "nameEN":"Tafalog",
      "nameCY":"Tafalog",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4031",
      "ngr":"SO2767029440",
      "titleEn":"Honddu at Tafalog",
      "titleCy":"Afon Honddu yn Nhafalog",
      "parameters":[
         {
            "parameter":81,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.433,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.030516,
         "longitude":-3.2169274
      },
      "location":4032,
      "nameEN":"Three Cocks",
      "nameCY":"Three Cocks",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4032",
      "ngr":"SO1661037600",
      "titleEn":"Llynfi at Three Cocks",
      "titleCy":"Afon Llynfi yn Three Cocks",
      "parameters":[
         {
            "parameter":93,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.626,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.745177,
         "longitude":-2.94852
      },
      "location":4039,
      "nameEN":"Chainbridge",
      "nameCY":"Pont Gadwyni",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4039",
      "ngr":"SO3461005590",
      "titleEn":"Usk at Chainbridge",
      "titleCy":"Afon Wysg ym Mhont Gadwyni",
      "parameters":[
         {
            "parameter":176,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":1.245,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.856798,
         "longitude":-3.1414316
      },
      "location":4040,
      "nameEN":"Crickhowell",
      "nameCY":"Crucywel",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4040",
      "ngr":"SO2148618195",
      "titleEn":"Usk at Crickhowell",
      "titleCy":"Afon Wysg yng Nghrucywel",
      "parameters":[
         {
            "parameter":177,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.512,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.875847,
         "longitude":-3.2698363
      },
      "location":4041,
      "nameEN":"Llandetty",
      "nameCY":"Llanddeti",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4041",
      "ngr":"SO1268020460",
      "titleEn":"Usk at Llandetty",
      "titleCy":"Afon Wysg yn Llanddeti",
      "parameters":[
         {
            "parameter":178,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.193,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.812118,
         "longitude":-3.0182482
      },
      "location":4042,
      "nameEN":"Llanfoist Bridge",
      "nameCY":"Pont Llan-ffwyst",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4042",
      "ngr":"SO2990013100",
      "titleEn":"Usk at Llanfoist Bridge",
      "titleCy":"Afon Wysg ym Mhont Llan-ffwyst",
      "parameters":[
         {
            "parameter":179,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.414,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.851535,
         "longitude":-3.1040723
      },
      "location":4043,
      "nameEN":"Millbrook",
      "nameCY":"Millbrook",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4043",
      "ngr":"SO2405017570",
      "titleEn":"Grwyne Fawr at Millbrook",
      "titleCy":"Afon Grwyne Fawr ym Millbrook",
      "parameters":[
         {
            "parameter":73,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.373,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.705162,
         "longitude":-2.8915327
      },
      "location":4044,
      "nameEN":"Olway Inn",
      "nameCY":"Olway Inn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4044",
      "ngr":"SO3849001090",
      "titleEn":"Olway Brook at Olway Inn",
      "titleCy":"Nant Olwy yn Olway Inn",
      "parameters":[
         {
            "parameter":114,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.219,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.963118,
         "longitude":-3.4517664
      },
      "location":4045,
      "nameEN":"Pont Ar Yscir",
      "nameCY":"Pont-ar-ysgir",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4045",
      "ngr":"SO0035030400",
      "titleEn":"Yscir at Pont Ar Yscir",
      "titleCy":"Afon Ysgir ym Mhont-ar-ysgir",
      "parameters":[
         {
            "parameter":203,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.523,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.917136,
         "longitude":-3.5599175
      },
      "location":4046,
      "nameEN":"Pont Hen Hafod",
      "nameCY":"Pont Hen Hafod",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4046",
      "ngr":"SN9281025440",
      "titleEn":"Senni at Pont Hen Hafod",
      "titleCy":"Afon Senni ym Mhont Hen Hafod",
      "parameters":[
         {
            "parameter":131,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.423,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.625782,
         "longitude":-2.9657612
      },
      "location":4047,
      "nameEN":"Ponthir",
      "nameCY":"Pont-hir",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4047",
      "ngr":"ST3324492327",
      "titleEn":"Lwyd at Ponthir",
      "titleCy":"Afon Lwyd ym Mhont-hir",
      "parameters":[
         {
            "parameter":100,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.604,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.607067,
         "longitude":-3.1028958
      },
      "location":4048,
      "nameEN":"Risca",
      "nameCY":"Rhisga",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4048",
      "ngr":"ST2372090380",
      "titleEn":"Ebbw at Risca",
      "titleCy":"Afon Ebwy yn Rhisga",
      "parameters":[
         {
            "parameter":54,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.640,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.881766,
         "longitude":-2.7886765
      },
      "location":4049,
      "nameEN":"Skenfrith",
      "nameCY":"Ynysgynwraidd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4049",
      "ngr":"SO4581020650",
      "titleEn":"Monnow at Skenfrith",
      "titleCy":"Afon Mynwy yn Ynysgynwraidd",
      "parameters":[
         {
            "parameter":107,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.938,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.954084,
         "longitude":-3.5331135
      },
      "location":4050,
      "nameEN":"Trallong",
      "nameCY":"Trallong",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4050",
      "ngr":"SN9474029510",
      "titleEn":"Usk at Trallong",
      "titleCy":"Llan-ffwyst yn Nhrallong",
      "parameters":[
         {
            "parameter":180,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.728,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.732824,
         "longitude":-2.9303055
      },
      "location":4051,
      "nameEN":"Trostrey Weir",
      "nameCY":"Cored Trostre",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4051",
      "ngr":"SO3585004200",
      "titleEn":"Usk at Trostrey Weir",
      "titleCy":"Afon Wysg yng Nghored Trostre",
      "parameters":[
         {
            "parameter":181,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.791,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.701973,
         "longitude":-2.9064471
      },
      "location":4054,
      "nameEN":"Usk Town",
      "nameCY":"Brynbuga",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4054",
      "ngr":"SO3745500748",
      "titleEn":"Usk at Usk Town",
      "titleCy":"Llan-ffwyst ym Mrynbuga",
      "parameters":[
         {
            "parameter":182,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.666,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.624492,
         "longitude":-3.1878286
      },
      "location":4055,
      "nameEN":"Ynysddu",
      "nameCY":"Ynys-ddu",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4055",
      "ngr":"ST1787092410",
      "titleEn":"Sirhowy at Ynysddu",
      "titleCy":"Afon Sirhywi yn Ynys-ddu",
      "parameters":[
         {
            "parameter":139,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.516,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.647758,
         "longitude":-3.3179566
      },
      "location":4056,
      "nameEN":"Fiddlers Elbow",
      "nameCY":"Fiddlers Elbow",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4056",
      "ngr":"ST0890995152",
      "titleEn":"Taff at Fiddlers Elbow",
      "titleCy":"Afon Taf yn Fiddlers Elbow",
      "parameters":[
         {
            "parameter":145,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.713,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.64412,
         "longitude":-3.4791874
      },
      "location":4057,
      "nameEN":"Gelli",
      "nameCY":"Y Gelli",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4057",
      "ngr":"SS9774694961",
      "titleEn":"Rhondda at Gelli",
      "titleCy":"Afon Rhondda yn y Gelli",
      "parameters":[
         {
            "parameter":119,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.293,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.534724,
         "longitude":-3.3946941
      },
      "location":4058,
      "nameEN":"Lanelay",
      "nameCY":"Glanelái",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4058",
      "ngr":"ST0336082680",
      "titleEn":"Ely at Lanelay",
      "titleCy":"Afon Elái yng Nglanelái",
      "parameters":[
         {
            "parameter":60,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.530,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.67639,
         "longitude":-3.4828567
      },
      "location":4059,
      "nameEN":"Maerdy",
      "nameCY":"Y Maerdy",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4059",
      "ngr":"SS9756598555",
      "titleEn":"Rhondda Fach at Maerdy",
      "titleCy":"Afon Rhondda Fach yn y Maerdy",
      "parameters":[
         {
            "parameter":123,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.405,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.598827,
         "longitude":-3.3305155
      },
      "location":4061,
      "nameEN":"Pontypridd",
      "nameCY":"Pontypridd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4061",
      "ngr":"ST0794189726",
      "titleEn":"Taff at Pontypridd",
      "titleCy":"Afon Taf ym Mhontypridd",
      "parameters":[
         {
            "parameter":146,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.939,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.485452,
         "longitude":-3.2676007
      },
      "location":4062,
      "nameEN":"St Fagans",
      "nameCY":"Sain Ffagan",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4062",
      "ngr":"ST1208077040",
      "titleEn":"Ely at St Fagans",
      "titleCy":"Afon Elái yn Sain Ffagan",
      "parameters":[
         {
            "parameter":61,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.690,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.609302,
         "longitude":-3.3688178
      },
      "location":4063,
      "nameEN":"Trehafod",
      "nameCY":"Trehafod",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4063",
      "ngr":"ST0531090940",
      "titleEn":"Rhondda at Trehafod",
      "titleCy":"Afon Rhondda yn Nhrehafod",
      "parameters":[
         {
            "parameter":120,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.523,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.708747,
         "longitude":-3.3459111
      },
      "location":4064,
      "nameEN":"Troedyrhiw",
      "nameCY":"Troed-y-rhiw",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4064",
      "ngr":"SO0710001970",
      "titleEn":"Taff at Troedyrhiw",
      "titleCy":"Afon Taf yn Nhroed-y-rhiw",
      "parameters":[
         {
            "parameter":147,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.718,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.676774,
         "longitude":-3.5451483
      },
      "location":4065,
      "nameEN":"Tynewydd",
      "nameCY":"Tynewydd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4065",
      "ngr":"SS9325998687",
      "titleEn":"Rhondda at Tynewydd",
      "titleCy":"Afon Rhondda yn Nhynewydd",
      "parameters":[
         {
            "parameter":121,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.354,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.497302,
         "longitude":-3.209873
      },
      "location":4067,
      "nameEN":"Western Avenue",
      "nameCY":"Rhodfa'r Gorllewin",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4067",
      "ngr":"ST1611078290",
      "titleEn":"Taff at Western Avenue",
      "titleCy":"Afon Taf ger Rhodfa'r Gorllewin",
      "parameters":[
         {
            "parameter":148,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.107,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.738084,
         "longitude":-3.5068175
      },
      "location":4068,
      "nameEN":"Hirwaun",
      "nameCY":"Hirwaun",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4068",
      "ngr":"SN9605005450",
      "titleEn":"Cynon at Hirwaun",
      "titleCy":"Afon Cynon yn Hirwaun",
      "parameters":[
         {
            "parameter":33,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.516,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.532726,
         "longitude":-3.1196958
      },
      "location":4070,
      "nameEN":"Llanedeyrn",
      "nameCY":"Llanedern",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4070",
      "ngr":"ST2243082130",
      "titleEn":"Rhymney at Llanedeyrn",
      "titleCy":"Afon Rhymni yn Llanedern",
      "parameters":[
         {
            "parameter":124,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.549,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.58877,
         "longitude":-3.1633676
      },
      "location":4071,
      "nameEN":"Waterloo Bridge, Machen",
      "nameCY":"Waterloo Bridge, Machen",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4071",
      "ngr":"ST1950088410",
      "titleEn":"Rhymney at Waterloo Bridge, Machen",
      "titleCy":"Afon Rhymni yn Waterloo Bridge, Machen",
      "parameters":[
         {
            "parameter":125,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.511,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.751781,
         "longitude":-3.3876092
      },
      "location":4072,
      "nameEN":"Merthyr Tydfil",
      "nameCY":"Merthyr Tudful",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4072",
      "ngr":"SO0431006810",
      "titleEn":"Taff at Merthyr Tydfil",
      "titleCy":"Afon Taf ym Merthyr Tudful",
      "parameters":[
         {
            "parameter":149,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.661,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.795113,
         "longitude":-3.2606062
      },
      "location":4073,
      "nameEN":"Nant Y Bwch",
      "nameCY":"Nant-y-Bwch",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4073",
      "ngr":"SO1316011470",
      "titleEn":"Sirhowy at Nant Y Bwch",
      "titleCy":"Afon Sirhywi yn Nant-y-Bwch",
      "parameters":[
         {
            "parameter":140,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.387,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.723775,
         "longitude":-3.2470382
      },
      "location":4074,
      "nameEN":"New Tredegar",
      "nameCY":"Tredegar Newydd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4074",
      "ngr":"SO1396003520",
      "titleEn":"Rhymney at New Tredegar",
      "titleCy":"Afon Rhymni yn Nhredegar Newydd",
      "parameters":[
         {
            "parameter":126,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.418,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.717802,
         "longitude":-3.0585419
      },
      "location":4075,
      "nameEN":"Pontnewynydd / Abersychan",
      "nameCY":"Pontnewynydd / Abersychan",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4075",
      "ngr":"SO2697002650",
      "titleEn":"Lwyd at Pontnewynydd / Abersychan",
      "titleCy":"Afon Lwyd ym Mhontnewynydd / Abersychan",
      "parameters":[
         {
            "parameter":101,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.325,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.593782,
         "longitude":-3.0713913
      },
      "location":4076,
      "nameEN":"Rhiwderin",
      "nameCY":"Rhiwderyn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4076",
      "ngr":"ST2588088870",
      "titleEn":"Ebbw at Rhiwderin",
      "titleCy":"Afon Ebwy yn Rhiwderyn",
      "parameters":[
         {
            "parameter":55,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.627,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.237863,
         "longitude":-4.257258
      },
      "location":4078,
      "nameEN":"Aberaeron",
      "nameCY":"Aberaeron",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4078",
      "ngr":"SN4596362365",
      "titleEn":"Aeron at Aberaeron",
      "titleCy":"Afon Aeron yn Aberaeron",
      "parameters":[
         {
            "parameter":1,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.075,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.5468,
         "longitude":-3.5918089
      },
      "location":4079,
      "nameEN":"Aberkenfig",
      "nameCY":"Abercynffig",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4079",
      "ngr":"SS8971884302",
      "titleEn":"Llynfi at Aberkenfig",
      "titleCy":"Afon Llynfi yn Abercynffig",
      "parameters":[
         {
            "parameter":94,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.487,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.41151,
         "longitude":-4.0863323
      },
      "location":4080,
      "nameEN":"Aberystwyth Tidal",
      "nameCY":"Llanw Aberystwyth ",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4080",
      "ngr":"SN5819081327",
      "titleEn":"Rheidol at Aberystwyth Tidal",
      "titleCy":"Llanw Afon Rheidol yn Aberystwyth ",
      "parameters":[
         {
            "parameter":116,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.831,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.598169,
         "longitude":-3.9950182
      },
      "location":4081,
      "nameEN":"Blackpill",
      "nameCY":"Dulais",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4081",
      "ngr":"SS6191790700",
      "titleEn":"Clyne at Blackpill",
      "titleCy":"Afon Clun yn Nulais",
      "parameters":[
         {
            "parameter":25,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.751,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.407869,
         "longitude":-3.4614125
      },
      "location":4082,
      "nameEN":"Upper Boverton",
      "nameCY":"Trebeferad Uchaf",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4082",
      "ngr":"SS9845168663",
      "titleEn":"Boverton Brook at Upper Boverton",
      "titleCy":"Nant Trebeferad yn Nhrebeferad Uchaf",
      "parameters":[
         {
            "parameter":12,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.743,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.503202,
         "longitude":-3.5807649
      },
      "location":4083,
      "nameEN":"Bridgend",
      "nameCY":"Pen-y-bont ar Ogwr",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4083",
      "ngr":"SS9037979437",
      "titleEn":"Ogmore at Bridgend",
      "titleCy":"Afon Ogwr ym Mhen-y-bont ar Ogwr",
      "parameters":[
         {
            "parameter":110,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.078,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.548032,
         "longitude":-3.5824772
      },
      "location":4084,
      "nameEN":"Brynmenyn",
      "nameCY":"Brynmenyn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4084",
      "ngr":"SS9036884425",
      "titleEn":"Ogmore at Brynmenyn",
      "titleCy":"Afon Ogwr ym Mrynmenyn",
      "parameters":[
         {
            "parameter":111,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.591,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.802951,
         "longitude":-4.7973706
      },
      "location":4085,
      "nameEN":"Canaston Bridge",
      "nameCY":"Pont Canaston",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4085",
      "ngr":"SN0722615292",
      "titleEn":"Eastern Cleddau at Canaston Bridge",
      "titleCy":"Afon Cleddau Ddu ym Mhont Canaston",
      "parameters":[
         {
            "parameter":53,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.093,
            "latestTime":"2020-10-06T21:30:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.692271,
         "longitude":-3.7686214
      },
      "location":4087,
      "nameEN":"Cilfrew",
      "nameCY":"Cil-ffriw",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4087",
      "ngr":"SN7785000761",
      "titleEn":"Dulais at Cilfrew",
      "titleCy":"Afon Dulais yng Nghil-ffriw",
      "parameters":[
         {
            "parameter":46,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.479,
            "latestTime":"2020-10-06T21:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.976784,
         "longitude":-4.9061292
      },
      "location":4088,
      "nameEN":"Cilrhedyn Bridge",
      "nameCY":"Pont Cilrhedyn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4088",
      "ngr":"SN0050134914",
      "titleEn":"Gwaun at Cilrhedyn Bridge",
      "titleCy":"Afon Gwaun ym Mhont Cilrhedyn",
      "parameters":[
         {
            "parameter":74,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.442,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.815096,
         "longitude":-4.5571824
      },
      "location":4089,
      "nameEN":"Clog Y Fran",
      "nameCY":"Clog y Fran",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4089",
      "ngr":"SN2383116034",
      "titleEn":"Taf at Clog Y Fran",
      "titleCy":"Afon Taf yng Nghlog y Fran",
      "parameters":[
         {
            "parameter":143,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.797,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.557244,
         "longitude":-3.6014347
      },
      "location":4090,
      "nameEN":"Coytrahen",
      "nameCY":"Goetre-hen",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4090",
      "ngr":"SS8907685478",
      "titleEn":"Llynfi at Coytrahen",
      "titleCy":"Afon Llynfi yng Nghoetre-hen",
      "parameters":[
         {
            "parameter":95,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.575,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.826146,
         "longitude":-3.6833826
      },
      "location":4091,
      "nameEN":"Craig y Nos",
      "nameCY":"Craig y Nos",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4091",
      "ngr":"SN8408515510",
      "titleEn":"Tawe at Craig y Nos",
      "titleCy":"Afon Tawe yn Nghraig y Nos",
      "parameters":[
         {
            "parameter":152,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.757,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.622579,
         "longitude":-3.7477213
      },
      "location":4092,
      "nameEN":"Cwmafan",
      "nameCY":"Cwmafan",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4092",
      "ngr":"SS7910992976",
      "titleEn":"Afan at Cwmafan",
      "titleCy":"Afon Afan yng Nghwmafan",
      "parameters":[
         {
            "parameter":3,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.390,
            "latestTime":"2020-10-06T21:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.39699,
         "longitude":-3.9224909
      },
      "location":4093,
      "nameEN":"Cwm Rheidol",
      "nameCY":"Cwm Rheidol",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4093",
      "ngr":"SN6929079403",
      "titleEn":"Rheidol at Cwm Rheidol",
      "titleCy":"Afon Rheidol yng Nghwm Rheidol",
      "parameters":[
         {
            "parameter":117,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.094,
            "latestTime":"2020-10-06T21:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.347995,
         "longitude":-3.7766991
      },
      "location":4094,
      "nameEN":"Cwmystwyth",
      "nameCY":"Cwmystwyth",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4094",
      "ngr":"SN7907573700",
      "titleEn":"Ystwyth at Cwmystwyth",
      "titleCy":"Afon Ystwyth yng Nghwmystwyth",
      "parameters":[
         {
            "parameter":204,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.750,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.520786,
         "longitude":-3.4961589
      },
      "location":4095,
      "nameEN":"Felindre Road",
      "nameCY":"Ffordd Felindre",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4095",
      "ngr":"SS9629181269",
      "titleEn":"Ewenny Fawr at Felindre Road",
      "titleCy":"Afon Ewenni Fawr yn Ffordd Felindre",
      "parameters":[
         {
            "parameter":65,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.747,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.829767,
         "longitude":-4.4832499
      },
      "location":4096,
      "nameEN":"Glasfryn Ford",
      "nameCY":"Rhyd Glasfryn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4096",
      "ngr":"SN2898217489",
      "titleEn":"Dewi Fawr at Glasfryn Ford",
      "titleCy":"Afon Dewi Fawr yn Rhyd Glasfryn",
      "parameters":[
         {
            "parameter":44,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.782,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.768833,
         "longitude":-3.7802803
      },
      "location":4097,
      "nameEN":"Gurnos",
      "nameCY":"Y Gurnos",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4097",
      "ngr":"SN7725209295",
      "titleEn":"Twrch at Gurnos",
      "titleCy":"Afon Twrch yn y Gurnos",
      "parameters":[
         {
            "parameter":165,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.681,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.804475,
         "longitude":-4.9701601
      },
      "location":4098,
      "nameEN":"Haverfordwest",
      "nameCY":"Hwlffordd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4098",
      "ngr":"SM9532215933",
      "titleEn":"Western Cleddau at Haverfordwest",
      "titleCy":"Afon Cleddau Wen yn Hwlffordd",
      "parameters":[
         {
            "parameter":188,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.694,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.492547,
         "longitude":-3.5571753
      },
      "location":4099,
      "nameEN":"Keepers Lodge",
      "nameCY":"Keepers Lodge",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4099",
      "ngr":"SS9199178217",
      "titleEn":"Ewenny at Keepers Lodge",
      "titleCy":"Afon Ewenni yn Keepers Lodge",
      "parameters":[
         {
            "parameter":64,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.544,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.671617,
         "longitude":-4.7123157
      },
      "location":4100,
      "nameEN":"Kiln Park",
      "nameCY":"Parc Kiln",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4100",
      "ngr":"SN1254600467",
      "titleEn":"Ritec at Kiln Park",
      "titleCy":"Afon Ritec ym Mharc Kiln",
      "parameters":[
         {
            "parameter":128,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.836,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.666436,
         "longitude":-4.1647964
      },
      "location":4102,
      "nameEN":"Llanelli Tidal",
      "nameCY":"Llanw Llanelli ",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4102",
      "ngr":"SS5038498626",
      "titleEn":"Llanelli at Llanelli Tidal",
      "titleCy":"Llanw Llanelli yn Llanelli ",
      "parameters":[
         {
            "parameter":97,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":3.264,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.800654,
         "longitude":-4.2423076
      },
      "location":4103,
      "nameEN":"Llangendeirne",
      "nameCY":"Llangyndeyrn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4103",
      "ngr":"SN4548313713",
      "titleEn":"Gwendraeth Fach at Llangendeirne",
      "titleCy":"Afon Gwendraeth Fach yn Llangyndeyrn",
      "parameters":[
         {
            "parameter":75,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.692,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.434044,
         "longitude":-4.0565474
      },
      "location":4104,
      "nameEN":"Llangorwen",
      "nameCY":"Llangorwen",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4104",
      "ngr":"SN6028783775",
      "titleEn":"Clarach at Llangorwen",
      "titleCy":"Afon Clarach yn Llangorwen",
      "parameters":[
         {
            "parameter":19,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.642,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.306747,
         "longitude":-4.1401657
      },
      "location":4105,
      "nameEN":"Llanrhystud",
      "nameCY":"Llanrhystud",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4105",
      "ngr":"SN5418469783",
      "titleEn":"Wyre at Llanrhystud",
      "titleCy":"Afon Wyre yn Llanrhystud",
      "parameters":[
         {
            "parameter":202,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.431,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.879059,
         "longitude":-4.6659264
      },
      "location":4106,
      "nameEN":"Login",
      "nameCY":"Login",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4106",
      "ngr":"SN1659723415",
      "titleEn":"Taf at Login",
      "titleCy":"Afon Taf yn Login",
      "parameters":[
         {
            "parameter":144,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.246,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.605807,
         "longitude":-3.6458885
      },
      "location":4107,
      "nameEN":"Maesteg",
      "nameCY":"Maesteg",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4107",
      "ngr":"SS8611690947",
      "titleEn":"Llynfi at Maesteg",
      "titleCy":"Afon Llynfi ym Maesteg",
      "parameters":[
         {
            "parameter":96,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.349,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.604175,
         "longitude":-3.7756791
      },
      "location":4108,
      "nameEN":"Marcroft Weir",
      "nameCY":"Cored Marcroft",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4108",
      "ngr":"SS7712490976",
      "titleEn":"Afan at Marcroft Weir",
      "titleCy":"Afon Afan yn Nghored Marcroft",
      "parameters":[
         {
            "parameter":4,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.974,
            "latestTime":"2020-10-06T21:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.791496,
         "longitude":-4.9786904
      },
      "location":4109,
      "nameEN":"Merlins Bridge",
      "nameCY":"Pont Fadlen",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4109",
      "ngr":"SM9467514514",
      "titleEn":"Merlin's Brook at Merlins Bridge",
      "titleCy":"Merlin’s Brook ym Mhont Fadlen",
      "parameters":[
         {
            "parameter":104,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.551,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.886191,
         "longitude":-5.1892605
      },
      "location":4110,
      "nameEN":"Middle Mill",
      "nameCY":"Felinganol",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4110",
      "ngr":"SM8061825657",
      "titleEn":"Solfach at Middle Mill",
      "titleCy":"Afon Solfach yn Felinganol",
      "parameters":[
         {
            "parameter":141,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.278,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.642181,
         "longitude":-3.9325036
      },
      "location":4111,
      "nameEN":"Morfa",
      "nameCY":"Afon Morfa",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4111",
      "ngr":"SS6637695478",
      "titleEn":"Tawe at Morfa",
      "titleCy":"Afon Tawe ym Morfa",
      "parameters":[
         {
            "parameter":153,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":3.522,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.664935,
         "longitude":-3.8087096
      },
      "location":4112,
      "nameEN":"Neath Tidal",
      "nameCY":"Llanw Nedd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4112",
      "ngr":"SS7500497789",
      "titleEn":"Neath at Neath Tidal",
      "titleCy":"Llanw Afon Nedd yn Nedd",
      "parameters":[
         {
            "parameter":108,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":3.635,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.603292,
         "longitude":-3.5409914
      },
      "location":4113,
      "nameEN":"Ogmore Vale",
      "nameCY":"Cwm Ogwr",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4113",
      "ngr":"SS9337490509",
      "titleEn":"Ogwr Fawr at Ogmore Vale",
      "titleCy":"Afon Ogwr Fawr yng Nghwm Ogwr",
      "parameters":[
         {
            "parameter":113,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.308,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.522064,
         "longitude":-3.4840352
      },
      "location":4114,
      "nameEN":"Old Mill",
      "nameCY":"Old Mill",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4114",
      "ngr":"SS9713581394",
      "titleEn":"Ewnenny Fach at Old Mill",
      "titleCy":"Afon Ewnenni Fach yn Old Mill",
      "parameters":[
         {
            "parameter":66,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.341,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.476884,
         "longitude":-3.6188501
      },
      "location":4115,
      "nameEN":"Penybont",
      "nameCY":"Pen-y-bont",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4115",
      "ngr":"SS8767176568",
      "titleEn":"Ogmore at Penybont",
      "titleCy":"Afon Ogwr ym Mhen-y-bont",
      "parameters":[
         {
            "parameter":112,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.266,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.795374,
         "longitude":-3.9736645
      },
      "location":4116,
      "nameEN":"Pont Amman",
      "nameCY":"Pontaman",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4116",
      "ngr":"SN6398912590",
      "titleEn":"Amman at Pont Amman",
      "titleCy":"Afon Aman ym Mhontaman",
      "parameters":[
         {
            "parameter":9,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.821,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.375448,
         "longitude":-4.0723786
      },
      "location":4117,
      "nameEN":"Pont Llolwyn",
      "nameCY":"Pont Llolwyn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4117",
      "ngr":"SN5902477289",
      "titleEn":"Ystwyth at Pont Llolwyn",
      "titleCy":"Afon Ystwyth ym Mhont Llolwyn",
      "parameters":[
         {
            "parameter":205,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.947,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.898375,
         "longitude":-5.1724016
      },
      "location":4118,
      "nameEN":"Pont y Cerbyd",
      "nameCY":"Pont y Cerbyd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4118",
      "ngr":"SM8183726961",
      "titleEn":"Solfach at Pont y Cerbyd",
      "titleCy":"Afon Solfach ym Mhont y Cerbyd",
      "parameters":[
         {
            "parameter":142,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.489,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.718425,
         "longitude":-3.8488319
      },
      "location":4119,
      "nameEN":"Pontardawe",
      "nameCY":"Pontardawe",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4119",
      "ngr":"SN7238003807",
      "titleEn":"Upper Clydach at Pontardawe",
      "titleCy":"Afon Clydach Uchaf ym Mhontardawe",
      "parameters":[
         {
            "parameter":175,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.360,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.7156,
         "longitude":-4.0297739
      },
      "location":4120,
      "nameEN":"Pontardulais",
      "nameCY":"Pontarddulais",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4120",
      "ngr":"SN5987303825",
      "titleEn":"Dulais at Pontardulais",
      "titleCy":"Afon Dulais ym Mhontarddulais",
      "parameters":[
         {
            "parameter":47,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.263,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.761783,
         "longitude":-3.5736689
      },
      "location":4121,
      "nameEN":"Pontneddfechan",
      "nameCY":"Pontneddfechan",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4121",
      "ngr":"SN9149108183",
      "titleEn":"Mellte at Pontneddfechan",
      "titleCy":"Afon Mellte ym Mhontneddfechan",
      "parameters":[
         {
            "parameter":103,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.635,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.753956,
         "longitude":-4.2190412
      },
      "location":4122,
      "nameEN":"Pontyates",
      "nameCY":"Pont-iets",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4122",
      "ngr":"SN4692908471",
      "titleEn":"Gwendraeth Fawr at Pontyates",
      "titleCy":"Afon Gwendraeth Fawr ym Mhont-iets",
      "parameters":[
         {
            "parameter":76,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.418,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.775038,
         "longitude":-4.1842849
      },
      "location":4123,
      "nameEN":"Pontyberem",
      "nameCY":"Pontyberem",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4123",
      "ngr":"SN4939810743",
      "titleEn":"Gwendraeth Fawr at Pontyberem",
      "titleCy":"Afon Gwendraeth Fawr ym Mhontyberem",
      "parameters":[
         {
            "parameter":77,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.615,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.82037,
         "longitude":-4.9700452
      },
      "location":4124,
      "nameEN":"Prendergast Mill",
      "nameCY":"Melin Prendergast",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4124",
      "ngr":"SM9540217700",
      "titleEn":"Western Cleddau at Prendergast Mill",
      "titleCy":"Afon Cleddau Wen ym Melin Prendergast",
      "parameters":[
         {
            "parameter":189,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.279,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.701481,
         "longitude":-3.7157396
      },
      "location":4125,
      "nameEN":"Resolven",
      "nameCY":"Resolfen",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4125",
      "ngr":"SN8152901698",
      "titleEn":"Neath at Resolven",
      "titleCy":"Afon Nedd yn Resolfen",
      "parameters":[
         {
            "parameter":109,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.846,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.819985,
         "longitude":-4.4962942
      },
      "location":4126,
      "nameEN":"St Clears",
      "nameCY":"Sanclêr",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4126",
      "ngr":"SN2804616432",
      "titleEn":"Cynin at St Clears",
      "titleCy":"Afon Cynin yn Sanclêr",
      "parameters":[
         {
            "parameter":32,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.659,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.736709,
         "longitude":-4.6983973
      },
      "location":4127,
      "nameEN":"Stepaside",
      "nameCY":"Stepaside",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4127",
      "ngr":"SN1377607669",
      "titleEn":"Fords Lake at Stepaside",
      "titleCy":"Fords Lake yn Stepaside",
      "parameters":[
         {
            "parameter":68,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.336,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.18463,
         "longitude":-4.1306231
      },
      "location":4129,
      "nameEN":"Talsarn",
      "nameCY":"Tal-sarn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4129",
      "ngr":"SN5443556183",
      "titleEn":"Aeron at Talsarn",
      "titleCy":"Afon Aeron yn Nhal-sarn",
      "parameters":[
         {
            "parameter":2,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.112,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.795704,
         "longitude":-3.9977381
      },
      "location":4131,
      "nameEN":"Tir Y Dail",
      "nameCY":"Tir y Dail",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4131",
      "ngr":"SN6233012672",
      "titleEn":"Loughor at Tir Y Dail",
      "titleCy":"Afon Llwchwr yn Nhir y Dail",
      "parameters":[
         {
            "parameter":98,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.868,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.818998,
         "longitude":-4.6096
      },
      "location":4132,
      "nameEN":"Whitland",
      "nameCY":"Hendy-Gwyn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4132",
      "ngr":"SN2023416596",
      "titleEn":"Gronw at Whitland",
      "titleCy":"Afon Gronw yn Hendy-gwyn ar Daf",
      "parameters":[
         {
            "parameter":72,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.563,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.712221,
         "longitude":-3.6956351
      },
      "location":4133,
      "nameEN":"Ynys Fach Bridge",
      "nameCY":"Pont Ynys Fach",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4133",
      "ngr":"SN8294602860",
      "titleEn":"Clydach Brook at Ynys Fach Bridge",
      "titleCy":"Nant Clydach ym Mhont Ynys Fach",
      "parameters":[
         {
            "parameter":24,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.341,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.681515,
         "longitude":-3.9026869
      },
      "location":4134,
      "nameEN":"Ynystanglws",
      "nameCY":"Ynystanglws",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4134",
      "ngr":"SS6855399798",
      "titleEn":"Tawe at Ynystanglws",
      "titleCy":"Afon Tawe yn Ynystanglws",
      "parameters":[
         {
            "parameter":154,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.933,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.731453,
         "longitude":-3.7499996
      },
      "location":4143,
      "nameEN":"Llawr Cae",
      "nameCY":"Llawr Cae",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4143",
      "ngr":"SH8192816305",
      "titleEn":"Cerist at Llawr Cae",
      "titleCy":"Afon Cerist yn Llawr Cae",
      "parameters":[
         {
            "parameter":18,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.219,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.031358,
         "longitude":-4.1266087
      },
      "location":4144,
      "nameEN":"Hafod Wydr",
      "nameCY":"Hafod Wydr",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4144",
      "ngr":"SH5748850346",
      "titleEn":"Colwyn at Hafod Wydr",
      "titleCy":"Afon Colwyn yn Hafod Wydr",
      "parameters":[
         {
            "parameter":27,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.286,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.105483,
         "longitude":-3.7921897
      },
      "location":4145,
      "nameEN":"Cwmlanerch",
      "nameCY":"Cwmlanerch",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4145",
      "ngr":"SH8011857978",
      "titleEn":"Conwy at Cwmlanerch",
      "titleCy":"Afon Conwy yng Nghwm Llannerch",
      "parameters":[
         {
            "parameter":28,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.788,
            "latestTime":"2020-10-06T21:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.600602,
         "longitude":-3.8550832
      },
      "location":4146,
      "nameEN":"Dyfi Bridge",
      "nameCY":"Pont Dyfi",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4146",
      "ngr":"SH7445801929",
      "titleEn":"Dyfi at Dyfi Bridge",
      "titleCy":"Afon Dyfi ym Mhont Dyfi",
      "parameters":[
         {
            "parameter":51,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":2.362,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.473289,
         "longitude":-4.0117456
      },
      "location":4147,
      "nameEN":"Dolybont",
      "nameCY":"Dol y bont",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4147",
      "ngr":"SN6345488054",
      "titleEn":"Leri at Dolybont",
      "titleCy":"Afon Tyleri yn Nôl-y-bont",
      "parameters":[
         {
            "parameter":87,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.398,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.067667,
         "longitude":-3.8131419
      },
      "location":4148,
      "nameEN":"Pont Gethin",
      "nameCY":"Pont Gethin",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4148",
      "ngr":"SH7860953807",
      "titleEn":"Lledr at Pont Gethin",
      "titleCy":"Afon Lledr ym Mhont Gethin",
      "parameters":[
         {
            "parameter":89,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.089,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.04567,
         "longitude":-4.3282907
      },
      "location":4149,
      "nameEN":"Pont y Cim",
      "nameCY":"Pont y Cim",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4149",
      "ngr":"SH4401752358",
      "titleEn":"Llyfni at Pont y Cim",
      "titleCy":"Afon Llyfni ym Mhont y Cim",
      "parameters":[
         {
            "parameter":92,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.437,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.82019,
         "longitude":-3.8770838
      },
      "location":4150,
      "nameEN":"Tyddyn Gwladys",
      "nameCY":"Tyddyn Gwladys",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4150",
      "ngr":"SH7360526391",
      "titleEn":"Mawddach at Tyddyn Gwladys",
      "titleCy":"Afon Mawddach yn Nhyddyn Gwladys",
      "parameters":[
         {
            "parameter":102,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.741,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.744403,
         "longitude":-3.8841207
      },
      "location":4151,
      "nameEN":"Dolgellau",
      "nameCY":"Dolgellau",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4151",
      "ngr":"SH7291017974",
      "titleEn":"Wnion at Dolgellau",
      "titleCy":"Afon Wnion yn Nolgellau",
      "parameters":[
         {
            "parameter":191,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":2.109,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.908556,
         "longitude":-3.5761997
      },
      "location":4152,
      "nameEN":"Bala",
      "nameCY":"Y Bala",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4152",
      "ngr":"SH9409535733",
      "titleEn":"Dee at Bala",
      "titleCy":"Afon Ddyfrdwy yn y Bala",
      "parameters":[
         {
            "parameter":37,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.611,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.910425,
         "longitude":-3.5897706
      },
      "location":4153,
      "nameEN":"Bala Weir X",
      "nameCY":"Gored X Y Bala",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4153",
      "ngr":"SH9318735961",
      "titleEn":"Tryweryn at Bala Weir X",
      "titleCy":"Afon Tryweryn yng Nghored X y Bala",
      "parameters":[
         {
            "parameter":164,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.661,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.008685,
         "longitude":-4.1002424
      },
      "location":4155,
      "nameEN":"Beddgelert",
      "nameCY":"Beddgelert",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4155",
      "ngr":"SH5918247772",
      "titleEn":"Glaslyn at Beddgelert",
      "titleCy":"Afon Glaslyn ym Meddgelert",
      "parameters":[
         {
            "parameter":71,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.018,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.231404,
         "longitude":-3.3420771
      },
      "location":4156,
      "nameEN":"Bodfari",
      "nameCY":"Bodfari",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4156",
      "ngr":"SJ1051271326",
      "titleEn":"Wheeler at Bodfari",
      "titleCy":"Afon Chwiler ym Modfari",
      "parameters":[
         {
            "parameter":190,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.508,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.265592,
         "longitude":-4.3565279
      },
      "location":4157,
      "nameEN":"Bodffordd",
      "nameCY":"Bodfford",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4157",
      "ngr":"SH4293076880",
      "titleEn":"Frogwy at Bodffordd",
      "titleCy":"Afon Frogwy ym Modffordd",
      "parameters":[
         {
            "parameter":69,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.808,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.114603,
         "longitude":-4.2669071
      },
      "location":4158,
      "nameEN":"Bontnewydd",
      "nameCY":"Bontnewydd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4158",
      "ngr":"SH4837459893",
      "titleEn":"Gwyrfai at Bontnewydd",
      "titleCy":"Afon Gwyrfai yn y Bontnewydd",
      "parameters":[
         {
            "parameter":80,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.694,
            "latestTime":"2020-10-06T21:30:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.928711,
         "longitude":-3.050064
      },
      "location":4162,
      "nameEN":"Brynkinalt",
      "nameCY":"Brynkinalt",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4162",
      "ngr":"SJ2951037328",
      "titleEn":"Ceiriog at Brynkinalt",
      "titleCy":"Afon Ceiriog ym Mryncunallt",
      "parameters":[
         {
            "parameter":17,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.703,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.979148,
         "longitude":-3.3879242
      },
      "location":4163,
      "nameEN":"Corwen",
      "nameCY":"Corwen",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4163",
      "ngr":"SJ0690843324",
      "titleEn":"Dee at Corwen",
      "titleCy":"Afon Ddyfrdwy yng Nghorwen",
      "parameters":[
         {
            "parameter":38,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.866,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.963272,
         "longitude":-3.7248956
      },
      "location":4164,
      "nameEN":"Cynefail",
      "nameCY":"Cynefail",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4164",
      "ngr":"SH8424242049",
      "titleEn":"River Gelyn Level at Cynefail",
      "titleCy":"Lefel Afon Gelyn yng Nghynefail",
      "parameters":[
         {
            "parameter":70,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.454,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.903094,
         "longitude":-3.5926989
      },
      "location":4165,
      "nameEN":"Dee Bridge",
      "nameCY":"Pont y Dee",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4165",
      "ngr":"SH9297235150",
      "titleEn":"Dee at Dee Bridge",
      "titleCy":"Afon Ddyfrdwy ym Mhont Dyfrdwy",
      "parameters":[
         {
            "parameter":39,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.241,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.981255,
         "longitude":-3.430145
      },
      "location":4166,
      "nameEN":"Druid",
      "nameCY":"Druid",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4166",
      "ngr":"SJ0407843614",
      "titleEn":"Alwen at Druid",
      "titleCy":"Afon Alwen yn y Ddwyryd",
      "parameters":[
         {
            "parameter":6,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.746,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.639739,
         "longitude":-4.0242104
      },
      "location":4167,
      "nameEN":"Pont y Garth",
      "nameCY":"Pont y Garth",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4167",
      "ngr":"SH6312706590",
      "titleEn":"Dysynni at Pont y Garth",
      "titleCy":"Afon Dysynni ym Mhont y Garth",
      "parameters":[
         {
            "parameter":52,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.056,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.007745,
         "longitude":-2.8664166
      },
      "location":4168,
      "nameEN":"Emral Brook",
      "nameCY":"Emral Brook",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4168",
      "ngr":"SJ4196145955",
      "titleEn":"Emral Brook at Emral Brook",
      "titleCy":"Emral Brook yn Emral Brook",
      "parameters":[
         {
            "parameter":62,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.936851,
         "longitude":-4.3821492
      },
      "location":4169,
      "nameEN":"Pencaenewydd",
      "nameCY":"Pencaenewydd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4169",
      "ngr":"SH4000540374",
      "titleEn":"Erch at Pencaenewydd",
      "titleCy":"Afon Erch ym Mhencaenewydd",
      "parameters":[
         {
            "parameter":63,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.191,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.083091,
         "longitude":-2.8791264
      },
      "location":4170,
      "nameEN":"Farndon",
      "nameCY":"Farndon",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4170",
      "ngr":"SJ4121154347",
      "titleEn":"Dee at Farndon",
      "titleCy":"Afon Ddyfrdwy yn Rhedynfre",
      "parameters":[
         {
            "parameter":40,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":6.743,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.962445,
         "longitude":-4.2367037
      },
      "location":4171,
      "nameEN":"Garndolbenmaen",
      "nameCY":"Garndolbenmaen",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4171",
      "ngr":"SH4986742906",
      "titleEn":"Dwyfor at Garndolbenmaen",
      "titleCy":"Afon Dwyfor yng Ngarndolbenmaen",
      "parameters":[
         {
            "parameter":50,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.535,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.134144,
         "longitude":-2.8713646
      },
      "location":4173,
      "nameEN":"Ironbridge",
      "nameCY":"Ironbridge",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4173",
      "ngr":"SJ4180060020",
      "titleEn":"Dee at Ironbridge",
      "titleCy":"Afon Ddyfrdwy yn Ironbridge",
      "parameters":[
         {
            "parameter":41,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":5.216,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.120018,
         "longitude":-2.8470366
      },
      "location":4174,
      "nameEN":"Lea Hall",
      "nameCY":"Lea Hall",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4174",
      "ngr":"SJ4340958429",
      "titleEn":"Aldford brook at Lea Hall",
      "titleCy":"Aldford Brook yn Lea Hall",
      "parameters":[
         {
            "parameter":5,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.945,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.137107,
         "longitude":-3.7972137
      },
      "location":4175,
      "nameEN":"Pont Fawr",
      "nameCY":"Pont Fawr",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4175",
      "ngr":"SH7987061504",
      "titleEn":"Conwy at Pont Fawr",
      "titleCy":"Afon Conwy ym Mhont Fawr",
      "parameters":[
         {
            "parameter":29,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"mAOD",
            "latestValue":5.707,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.96665,
         "longitude":-2.9716221
      },
      "location":4176,
      "nameEN":"Manley Hall",
      "nameCY":"Manley Hall",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4176",
      "ngr":"SJ3484041474",
      "titleEn":"Dee at Manley Hall",
      "titleCy":"Afon Ddyfrdwy ym Manley Hall",
      "parameters":[
         {
            "parameter":42,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.115,
            "latestTime":"2020-10-06T21:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.100617,
         "longitude":-4.0801768
      },
      "location":4177,
      "nameEN":"Nant Peris",
      "nameCY":"Nant Peris",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4177",
      "ngr":"SH6082557958",
      "titleEn":"Peris at Nant Peris",
      "titleCy":"Afon Peris yn Nant Peris",
      "parameters":[
         {
            "parameter":115,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.010,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.862986,
         "longitude":-3.6745809
      },
      "location":4178,
      "nameEN":"New Inn",
      "nameCY":"New Inn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4178",
      "ngr":"SH8736130814",
      "titleEn":"Dee at New Inn",
      "titleCy":"Afon Ddyfrdwy yn New Inn",
      "parameters":[
         {
            "parameter":43,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.047,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.136374,
         "longitude":-4.2520264
      },
      "location":4179,
      "nameEN":"Peblic Mill",
      "nameCY":"Melin Peblig",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4179",
      "ngr":"SH4944662283",
      "titleEn":"Seiont at Peblic Mill",
      "titleCy":"Afon Seiont ym Melin Peblig",
      "parameters":[
         {
            "parameter":130,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.072,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.262457,
         "longitude":-3.4344599
      },
      "location":4180,
      "nameEN":"Pont Dafydd",
      "nameCY":"Pont Dafydd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4180",
      "ngr":"SJ0441574900",
      "titleEn":"Clwyd at Pont Dafydd",
      "titleCy":"Afon Clwyd ym Mhont Dafydd",
      "parameters":[
         {
            "parameter":20,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":2.797,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.227635,
         "longitude":-3.3947504
      },
      "location":4181,
      "nameEN":"Pont y Cambwll",
      "nameCY":"Pont y Cambwll",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4181",
      "ngr":"SJ0698870974",
      "titleEn":"Clwyd at Pont y Cambwll",
      "titleCy":"Afon Clwyd ym Mhont y Cambwll",
      "parameters":[
         {
            "parameter":21,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":1.656,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.079402,
         "longitude":-2.9937308
      },
      "location":4182,
      "nameEN":"Pont y Capel",
      "nameCY":"Pont y Capel",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4182",
      "ngr":"SJ3352954037",
      "titleEn":"Alyn at Pont y Capel",
      "titleCy":"Afon Alun ym Mhont y Capel",
      "parameters":[
         {
            "parameter":7,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.854,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.232715,
         "longitude":-3.5709448
      },
      "location":4183,
      "nameEN":"Pont y Gwyddel",
      "nameCY":"Pont y Gwyddel",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4183",
      "ngr":"SH9523971783",
      "titleEn":"Elwy at Pont y Gwyddel",
      "titleCy":"Afon Elwy ym Mhont y Gwyddel",
      "parameters":[
         {
            "parameter":58,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.211,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.293488,
         "longitude":-3.4750044
      },
      "location":4185,
      "nameEN":"Rhuddlan",
      "nameCY":"Rhuddlan",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4185",
      "ngr":"SJ0178278407",
      "titleEn":"Clwyd at Rhuddlan",
      "titleCy":"Afon Clwyd yn Rhuddlan",
      "parameters":[
         {
            "parameter":22,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":3.228,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.191847,
         "longitude":-3.1927941
      },
      "location":4186,
      "nameEN":"Rhydymwyn",
      "nameCY":"Rhydymwyn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4186",
      "ngr":"SJ2040366749",
      "titleEn":"Alyn at Rhydymwyn",
      "titleCy":"Afon Alun yn Rhyd-y-mwyn",
      "parameters":[
         {
            "parameter":8,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.059,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.122874,
         "longitude":-3.3137365
      },
      "location":4187,
      "nameEN":"Ruthin Weir",
      "nameCY":"Cored Rhuthun",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4187",
      "ngr":"SJ1218259218",
      "titleEn":"Clwyd at Ruthin Weir",
      "titleCy":"Afon Clwyd yng Nghored Rhuthun",
      "parameters":[
         {
            "parameter":23,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.642,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.262103,
         "longitude":-3.4495457
      },
      "location":4188,
      "nameEN":"St Asaph",
      "nameCY":"Llanelwy",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4188",
      "ngr":"SJ0340874881",
      "titleEn":"Elwy at St Asaph",
      "titleCy":"Afon Elwy yn Llanelwy",
      "parameters":[
         {
            "parameter":59,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":2.113,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.160207,
         "longitude":-3.8241158
      },
      "location":4191,
      "nameEN":"Trefriw",
      "nameCY":"Trefriw",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4191",
      "ngr":"SH7813664119",
      "titleEn":"Conwy at Trefriw",
      "titleCy":"Afon Conwy yn Nhrefriw",
      "parameters":[
         {
            "parameter":30,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":3.172,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.007827,
         "longitude":-2.8662841
      },
      "location":4192,
      "nameEN":"Wych Brook",
      "nameCY":"Nant Wych",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4192",
      "ngr":"SJ4197045964",
      "titleEn":"Wych Brook at Wych Brook",
      "titleCy":"Nant Wych yn Nant Wych",
      "parameters":[
         {
            "parameter":192,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":2.178,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.863141,
         "longitude":-4.201464
      },
      "location":4193,
      "nameEN":"Capel Dewi",
      "nameCY":"Capel Dewi",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4193",
      "ngr":"SN4850920576",
      "titleEn":"Tywi at Capel Dewi",
      "titleCy":"Afon Tywi yng Nghapel Dewi",
      "parameters":[
         {
            "parameter":168,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.862,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.109412,
         "longitude":-3.7924605
      },
      "location":4194,
      "nameEN":"Craig Clungwyn",
      "nameCY":"Craig Clungwyn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4194",
      "ngr":"SN7734547192",
      "titleEn":"Doethie at Craig Clungwyn",
      "titleCy":"Afon Doethie yng Nghraig Clungwyn",
      "parameters":[
         {
            "parameter":45,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.023,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.077768,
         "longitude":-3.971084
      },
      "location":4195,
      "nameEN":"Ddol Las",
      "nameCY":"Ddol Las",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4195",
      "ngr":"SN6501843990",
      "titleEn":"Twrch at Ddol Las",
      "titleCy":"Afon Twrch yn Nôl-Las",
      "parameters":[
         {
            "parameter":166,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.666,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.010773,
         "longitude":-3.8053884
      },
      "location":4196,
      "nameEN":"Dolau Hirion",
      "nameCY":"Dolau Hirion",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4196",
      "ngr":"SN7618736244",
      "titleEn":"Tywi at Dolau Hirion",
      "titleCy":"Afon Tywi yn Nolau Hirion",
      "parameters":[
         {
            "parameter":169,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.893,
            "latestTime":"2020-10-06T19:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.880974,
         "longitude":-4.1680298
      },
      "location":4197,
      "nameEN":"Felin Mynachdy",
      "nameCY":"Felin Mynachdy",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4197",
      "ngr":"SN5087022490",
      "titleEn":"Cothi at Felin Mynachdy",
      "titleCy":"Afon Cothi yn Felin Mynachdy",
      "parameters":[
         {
            "parameter":31,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.469,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.922738,
         "longitude":-3.8742037
      },
      "location":4198,
      "nameEN":"Felin Y Cwm",
      "nameCY":"Felin y Cwm",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4198",
      "ngr":"SN7121226573",
      "titleEn":"Sawdde at Felin Y Cwm",
      "titleCy":"Afon Sawdde yn Felin y Cwm",
      "parameters":[
         {
            "parameter":129,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.598,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.873944,
         "longitude":-4.2815034
      },
      "location":4199,
      "nameEN":"Glan Gwili",
      "nameCY":"Glan Gwili",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4199",
      "ngr":"SN4303621947",
      "titleEn":"Gwili at Glan Gwili",
      "titleCy":"Afon Gwili yng Nglan Gwili",
      "parameters":[
         {
            "parameter":78,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.929,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.993882,
         "longitude":-3.7909876
      },
      "location":4201,
      "nameEN":"Llandovery",
      "nameCY":"Llanymddyfri",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4201",
      "ngr":"SN7712934341",
      "titleEn":"Bran at Llandovery",
      "titleCy":"Afon Brân yn Llanymddyfri",
      "parameters":[
         {
            "parameter":13,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.758,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.94019,
         "longitude":-3.8843733
      },
      "location":4202,
      "nameEN":"Llangadog",
      "nameCY":"Llangadog",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4202",
      "ngr":"SN7056328532",
      "titleEn":"Bran at Llangadog",
      "titleCy":"Afon Brân yn Llangadog",
      "parameters":[
         {
            "parameter":14,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.759,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.967319,
         "longitude":-3.8725698
      },
      "location":4203,
      "nameEN":"Llanwrda",
      "nameCY":"Llanwrda",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4203",
      "ngr":"SN7145231528",
      "titleEn":"Dulais at Llanwrda",
      "titleCy":"Afon Dulais yn Llanwrda",
      "parameters":[
         {
            "parameter":48,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.523,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.121775,
         "longitude":-3.7644282
      },
      "location":4204,
      "nameEN":"Llyn Brianne Reservoir",
      "nameCY":"Argae Llyn Brianne",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4204",
      "ngr":"SN7929848520",
      "titleEn":"Tywi at Llyn Brianne Reservoir",
      "titleCy":"Afon Tywi yn Argae Llyn Brianne",
      "parameters":[
         {
            "parameter":170,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":-2.115,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.89827,
         "longitude":-3.952721
      },
      "location":4205,
      "nameEN":"Manorafon",
      "nameCY":"Manorafon",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4205",
      "ngr":"SN6574023994",
      "titleEn":"Tywi at Manorafon",
      "titleCy":"Afon Tywi ym Manorafon",
      "parameters":[
         {
            "parameter":171,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.891,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.99297,
         "longitude":-3.7799834
      },
      "location":4206,
      "nameEN":"Pont Felindre",
      "nameCY":"Pont Felindre",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4206",
      "ngr":"SN7788234221",
      "titleEn":"Gwydderig at Pont Felindre",
      "titleCy":"Afon Gwydderig ym Mhont Felindre",
      "parameters":[
         {
            "parameter":79,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.885,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.853952,
         "longitude":-4.3101593
      },
      "location":4207,
      "nameEN":"Pothouse Wharf",
      "nameCY":"Pothouse Wharf",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4207",
      "ngr":"SN4099319786",
      "titleEn":"Tywi at Pothouse Wharf",
      "titleCy":"Afon Tywi yn Pothouse Wharf",
      "parameters":[
         {
            "parameter":172,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":3.708,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.861705,
         "longitude":-4.1921711
      },
      "location":4208,
      "nameEN":"Ty Castell",
      "nameCY":"Ty Castell",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4208",
      "ngr":"SN4914420397",
      "titleEn":"Tywi at Ty Castell",
      "titleCy":"Afon Tywi yn Nhy Castell",
      "parameters":[
         {
            "parameter":173,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":2.014,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.11012,
         "longitude":-3.77476
      },
      "location":4209,
      "nameEN":"Ystradffin",
      "nameCY":"Ystradffin",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4209",
      "ngr":"SN7855947241",
      "titleEn":"Tywi at Ystradffin",
      "titleCy":"Afon Tywi yn Ystradffin",
      "parameters":[
         {
            "parameter":174,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.199,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.045405,
         "longitude":-4.5617405
      },
      "location":4210,
      "nameEN":"Glan Teifi",
      "nameCY":"Glan Teifi",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4210",
      "ngr":"SN2441941655",
      "titleEn":"Teifi at Glan Teifi",
      "titleCy":"Afon Teifi yng Nglan Teifi",
      "parameters":[
         {
            "parameter":156,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":2.887,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.036513,
         "longitude":-4.3173034
      },
      "location":4211,
      "nameEN":"Llandysul",
      "nameCY":"Llandysul",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4211",
      "ngr":"SN4114840103",
      "titleEn":"Tyweli at Llandysul",
      "titleCy":"Afon Tyweli yn Llandysul",
      "parameters":[
         {
            "parameter":167,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.988,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.042094,
         "longitude":-4.2857
      },
      "location":4212,
      "nameEN":"Llanfair",
      "nameCY":"Llanfair",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4212",
      "ngr":"SN4333540655",
      "titleEn":"Teifi at Llanfair",
      "titleCy":"Afon Teifi yn Llanfair",
      "parameters":[
         {
            "parameter":157,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.844,
            "latestTime":"2020-10-06T21:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.193728,
         "longitude":-3.9731095
      },
      "location":4213,
      "nameEN":"Pont Llanio",
      "nameCY":"Pont Llanio",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4213",
      "ngr":"SN6523056890",
      "titleEn":"Teifi at Pont Llanio",
      "titleCy":"Afon Teifi ym Mhont Llanio",
      "parameters":[
         {
            "parameter":158,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.985,
            "latestTime":"2020-10-06T21:45:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.28202,
         "longitude":-3.8627965
      },
      "location":4214,
      "nameEN":"Pontrhydfendigaid",
      "nameCY":"Pontrhydfendigaid",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4214",
      "ngr":"SN7302266510",
      "titleEn":"Teifi at Pontrhydfendigaid",
      "titleCy":"Afon Teifi ym Mhontrhydfendigaid",
      "parameters":[
         {
            "parameter":159,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.165,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.220373,
         "longitude":-3.9333274
      },
      "location":4216,
      "nameEN":"Tregaron",
      "nameCY":"Tregaron",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4216",
      "ngr":"SN6802859780",
      "titleEn":"Brennig at Tregaron",
      "titleCy":"Afon Brenig yn Nhregaron",
      "parameters":[
         {
            "parameter":15,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.757,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.462147,
         "longitude":-3.4460877
      },
      "location":4217,
      "nameEN":"Cowbridge",
      "nameCY":"Y Bont-faen",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4217",
      "ngr":"SS9963674678",
      "titleEn":"Thaw at Cowbridge",
      "titleCy":"Afon Ddawan yn y Bont-faen",
      "parameters":[
         {
            "parameter":161,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.498,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.435406,
         "longitude":-3.4150018
      },
      "location":4218,
      "nameEN":"Gigman Bridge",
      "nameCY":"Pont Gigman",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4218",
      "ngr":"ST0173871662",
      "titleEn":"Thaw at Gigman Bridge",
      "titleCy":"Afon Ddawan ym Mhont Gigman",
      "parameters":[
         {
            "parameter":162,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.447,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.574987,
         "longitude":-3.2920973
      },
      "location":4219,
      "nameEN":"Upper Boat Bridge",
      "nameCY":"Pont Glan-bad",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4219",
      "ngr":"ST1055587027",
      "titleEn":"Taff at Upper Boat Bridge",
      "titleCy":"Afon Taf ym Mhont Glan-bad",
      "parameters":[
         {
            "parameter":150,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.816,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.605117,
         "longitude":-3.3379485
      },
      "location":4220,
      "nameEN":"Pontypridd, Sion Street",
      "nameCY":"Pontypridd, Stryd Sion",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4220",
      "ngr":"ST0743990435",
      "titleEn":"Taff at Pontypridd, Sion Street",
      "titleCy":"Afon Taf ym Mhontypridd, Stryd Sion",
      "parameters":[
         {
            "parameter":151,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":4.392,
            "latestTime":"2020-02-16T04:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.652094,
         "longitude":-3.3320895
      },
      "location":4221,
      "nameEN":"Abercynon",
      "nameCY":"Abercynon",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4221",
      "ngr":"ST0794095652",
      "titleEn":"Cynon at Abercynon",
      "titleCy":"Afon Cynon yn Abercynon",
      "parameters":[
         {
            "parameter":34,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.518,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.677829,
         "longitude":-3.2221134
      },
      "location":4222,
      "nameEN":"Bargoed",
      "nameCY":"Bargoed",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4222",
      "ngr":"ST1559698381",
      "titleEn":"Rhymney at Bargoed",
      "titleCy":"Afon Rhymni ym Margod",
      "parameters":[
         {
            "parameter":127,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.505,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.081111,
         "longitude":-4.6617115
      },
      "location":4224,
      "nameEN":"Cardigan Quay",
      "nameCY":"Cei Aberteifi",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4224",
      "ngr":"SN1771045872",
      "titleEn":"Teifi at Cardigan Quay",
      "titleCy":"Afon Teifi yng Nghei Aberteifi",
      "parameters":[
         {
            "parameter":160,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.042,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.780424,
         "longitude":-3.7575561
      },
      "location":4225,
      "nameEN":"Teddy Bear Bridge",
      "nameCY":"Teddy Bear Bridge",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4225",
      "ngr":"SN7885110546",
      "titleEn":"Tawe at Teddy Bear Bridge",
      "titleCy":"Afon Tawe yn Teddy Bear Bridge",
      "parameters":[
         {
            "parameter":155,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.781,
            "latestTime":"2020-02-19T12:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.591334,
         "longitude":-3.7679743
      },
      "location":4226,
      "nameEN":"Taibach",
      "nameCY":"Tai-bach",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4226",
      "ngr":"SS7762389535",
      "titleEn":"Ffrwdwyllt at Taibach",
      "titleCy":"Afon Ffrwdwyllt yn Nhai-bach",
      "parameters":[
         {
            "parameter":67,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.387,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.699011,
         "longitude":-4.1454977
      },
      "location":4227,
      "nameEN":"Felinfoel",
      "nameCY":"Felin-foel",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4227",
      "ngr":"SN5182502209",
      "titleEn":"Lliedi at Felinfoel",
      "titleCy":"Afon Lliedi yn Felin-foel",
      "parameters":[
         {
            "parameter":90,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.457,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.682963,
         "longitude":-4.134887
      },
      "location":4228,
      "nameEN":"Halfway",
      "nameCY":"Halfway",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4228",
      "ngr":"SN5250600403",
      "titleEn":"Dafen at Halfway",
      "titleCy":"Afon Dafen yn Halfway",
      "parameters":[
         {
            "parameter":36,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.501,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.655123,
         "longitude":-4.0552328
      },
      "location":4229,
      "nameEN":"Pont-y-Cob",
      "nameCY":"Pont-y-Cob",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4229",
      "ngr":"SS5792597149",
      "titleEn":"Lliw at Pont-y-Cob",
      "titleCy":"Afon Lliw ym Mhont-y-Cob",
      "parameters":[
         {
            "parameter":91,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":3.339,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.403337,
         "longitude":-4.0576176
      },
      "location":4230,
      "nameEN":"Llanbadarn Fawr",
      "nameCY":"Llanbadarn Fawr",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4230",
      "ngr":"SN6011780362",
      "titleEn":"Rheidol at Llanbadarn Fawr",
      "titleCy":"Afon Rheidol yn Llanbadarn Fawr",
      "parameters":[
         {
            "parameter":118,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.097,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.713126,
         "longitude":-3.440677
      },
      "location":4231,
      "nameEN":"Aberdare",
      "nameCY":"Aberdâr",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4231",
      "ngr":"SO0056202582",
      "titleEn":"Cynon at Aberdare",
      "titleCy":"Afon Cynon yn Aberdâr",
      "parameters":[
         {
            "parameter":35,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.391,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.706354,
         "longitude":-3.1450356
      },
      "location":4232,
      "nameEN":"Aberbeeg",
      "nameCY":"Aber-bîg",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4232",
      "ngr":"SO2097501467",
      "titleEn":"Ebbw at Aberbeeg",
      "titleCy":"Afon Ebwy yn Aber-bîg",
      "parameters":[
         {
            "parameter":56,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.616,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.947681,
         "longitude":-3.3944346
      },
      "location":4233,
      "nameEN":"Brecon",
      "nameCY":"Aberhonddu",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4233",
      "ngr":"SO0425628606",
      "titleEn":"Usk at Brecon",
      "titleCy":"Afon Wysg yn Aberhonddu",
      "parameters":[
         {
            "parameter":183,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.787,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.515882,
         "longitude":-3.3121906
      },
      "location":4235,
      "nameEN":"Newtown",
      "nameCY":"Y Drenewydd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4235",
      "ngr":"SO1105391698",
      "titleEn":"Severn at Newtown",
      "titleCy":"Afon Hafren yn y Drenewydd",
      "parameters":[
         {
            "parameter":138,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.278,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.000857,
         "longitude":-2.9112057
      },
      "location":4241,
      "nameEN":"Bangor on Dee",
      "nameCY":"Bangor is y Coed",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4241",
      "ngr":"SJ3894645226",
      "titleEn":"Dee at Bangor on Dee",
      "titleCy":"Afon Dyfrdwy am Bont Leadmill, Y Wyddgrug",
      "parameters":[
         {
            "parameter":10062,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"mAOD",
            "latestValue":13.133,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.172387,
         "longitude":-3.1374823
      },
      "location":4242,
      "nameEN":"Mold Leadmill Bridge",
      "nameCY":"Bont Leadmill, Y Wyddgrug",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4242",
      "ngr":"SJ2406464524",
      "titleEn":"Alyn at Mold Leadmill Bridge",
      "titleCy":"Afon Alun ym Mhont Leadmill, Y Wyddgrug",
      "parameters":[
         {
            "parameter":10063,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.279,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.137365,
         "longitude":-3.0824575
      },
      "location":4243,
      "nameEN":"Pontblyddyn",
      "nameCY":"Pontblyddyn",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4243",
      "ngr":"SJ2768360571",
      "titleEn":"Alyn at Pontblyddyn",
      "titleCy":"Afon Alun ym Mhontblyddyn",
      "parameters":[
         {
            "parameter":10064,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.379,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.107304,
         "longitude":-2.9520844
      },
      "location":4244,
      "nameEN":"Rossett Alyn Bridge",
      "nameCY":"Bont Alyn, Yr Orsedd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4244",
      "ngr":"SJ3636057103",
      "titleEn":"Alyn at Rossett Alyn Bridge",
      "titleCy":"Afon Alun am Mhont Alyn, Yr Orsedd",
      "parameters":[
         {
            "parameter":10065,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.910,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.535996,
         "longitude":-3.2749431
      },
      "location":4259,
      "nameEN":"Aberbechan",
      "nameCY":"Aberbechan",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4259",
      "ngr":"SO1362093890",
      "titleEn":"Bechan Brook at Aberbechan",
      "titleCy":"Nant Bechan yn Aberbechan",
      "parameters":[
         {
            "parameter":10066,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.804,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.455202,
         "longitude":-3.4885196
      },
      "location":4260,
      "nameEN":"Dolwen",
      "nameCY":"Dolwen",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4260",
      "ngr":"SN9894985181",
      "titleEn":"River Severn at Dolwen",
      "titleCy":"Afon Hafren yn Dolwen",
      "parameters":[
         {
            "parameter":10067,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.959,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.763642,
         "longitude":-3.2645846
      },
      "location":4261,
      "nameEN":"Llanfyllin (Cain)",
      "nameCY":"Llanfyllin (Cain)",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4261",
      "ngr":"SJ1476719199",
      "titleEn":"Cain at Llanfyllin",
      "titleCy":"Afon Cain yn Llanfyllin",
      "parameters":[
         {
            "parameter":10068,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.380,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.775935,
         "longitude":-3.1741992
      },
      "location":4262,
      "nameEN":"Llansantffraid",
      "nameCY":"Llansantffraid",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4262",
      "ngr":"SJ2088820463",
      "titleEn":"Cain at Llansantffraid",
      "titleCy":"Afon Cain yn Llansantffraid",
      "parameters":[
         {
            "parameter":10069,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.977,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.591665,
         "longitude":-3.1613089
      },
      "location":4263,
      "nameEN":"Pont-y-Gaer",
      "nameCY":"Pont-y-Gaer",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4263",
      "ngr":"SO2142799952",
      "titleEn":"Camlad at Pont-y-Gaer",
      "titleCy":"Afon Camlad ym Mhont-y-Gaer",
      "parameters":[
         {
            "parameter":10070,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":1.145,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.879582,
         "longitude":-4.9669448
      },
      "location":4264,
      "nameEN":"Treffgarne",
      "nameCY":"Treffgarne",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4264",
      "ngr":"SM9588424275",
      "titleEn":"Western Cleddau at Treffgarne",
      "titleCy":"Afon Cleddau orllewin yn Nhrefgarn",
      "parameters":[
         {
            "parameter":10071,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.610,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.593017,
         "longitude":-3.7909658
      },
      "location":4265,
      "nameEN":"Green Park Weir",
      "nameCY":"Cored Green Park",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4265",
      "ngr":"SS7603589761",
      "titleEn":"River Afan at Greenpark Weir",
      "titleCy":"Afon Afan yng Nghored Green Park",
      "parameters":[
         {
            "parameter":10072,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":3.716,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.833014,
         "longitude":-3.9527407
      },
      "location":4266,
      "nameEN":"Forge Bridge",
      "nameCY":"Pont Forge",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4266",
      "ngr":"SN6554416737",
      "titleEn":"Loughor at Forge Bridge",
      "titleCy":"Llwchwr ym Mhont Forge",
      "parameters":[
         {
            "parameter":10073,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.320,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.483186,
         "longitude":-3.9826098
      },
      "location":4267,
      "nameEN":"Tal y Bont Leri",
      "nameCY":"Tal y Bont Leri",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4267",
      "ngr":"SN6546389100",
      "titleEn":"Leri at Tal y Bont",
      "titleCy":"Leri yn Nhal y Bont",
      "parameters":[
         {
            "parameter":10074,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.536,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.697373,
         "longitude":-4.1666203
      },
      "location":4268,
      "nameEN":"Trebeddrod",
      "nameCY":"Trebeddrod",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4268",
      "ngr":"SN5036002070",
      "titleEn":"Trebeddrod Reservoir Level",
      "titleCy":"Trebeddrod Cronfa Ddwr Lefel",
      "parameters":[
         {
            "parameter":10075,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.000,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.747541,
         "longitude":-3.6166175
      },
      "location":4269,
      "nameEN":"Glynneath",
      "nameCY":"Glyn-nedd",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4269",
      "ngr":"SN8849206664",
      "titleEn":"River Neath at Glynneath",
      "titleCy":"Afon Nedd yng Nglyn-nedd",
      "parameters":[
         {
            "parameter":10076,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.187,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.820865,
         "longitude":-4.0051343
      },
      "location":4270,
      "nameEN":"Llandybie",
      "nameCY":"Llandybie",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4270",
      "ngr":"SN6189715484",
      "titleEn":"River Marlas at Llandybie",
      "titleCy":"Afon Marlas yn Llandybie",
      "parameters":[
         {
            "parameter":10081,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.329,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.772413,
         "longitude":-3.7611148
      },
      "location":4271,
      "nameEN":"Ystradgynlais",
      "nameCY":"Ystradgynlais",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4271",
      "ngr":"SN7858409661",
      "titleEn":"Tawe at Ystradgynlais",
      "titleCy":"Afon Tawe yn Ystradgynlais",
      "parameters":[
         {
            "parameter":10077,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.873,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.697865,
         "longitude":-4.0812204
      },
      "location":4272,
      "nameEN":"Llangennech",
      "nameCY":"Llangennech",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4272",
      "ngr":"SN5626301953",
      "titleEn":"River Morlais at Llangennech",
      "titleCy":"Afon Morlais yn Llangennech",
      "parameters":[
         {
            "parameter":10078,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.546,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.689352,
         "longitude":-4.189358
      },
      "location":4273,
      "nameEN":"Stradey Home Farm",
      "nameCY":"Fferm garter y Strade ",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4273",
      "ngr":"SN4876201225",
      "titleEn":"Dulais at Stradey Home Farm",
      "titleCy":"Dulais yn Fferm Gartre'r Strade",
      "parameters":[
         {
            "parameter":10079,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.387,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.799216,
         "longitude":-4.5313643
      },
      "location":4274,
      "nameEN":"Llanddowror",
      "nameCY":"Llanddowror",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4274",
      "ngr":"SN2554914206",
      "titleEn":"Llanddowror Brook at Llanddowror",
      "titleCy":"Nant Llanddowror yn Llanddowror",
      "parameters":[
         {
            "parameter":10080,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.325,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.410578,
         "longitude":-3.4597878
      },
      "location":4275,
      "nameEN":"Llanmaes",
      "nameCY":"Llanmaes",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4275",
      "ngr":"SS9857068962",
      "titleEn":"Llanmaes Brook at Llanmaes",
      "titleCy":"Nant Llanmaes yn Llanmaes",
      "parameters":[
         {
            "parameter":10082,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.797,
            "latestTime":"2020-10-06T20:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.809104,
         "longitude":-3.4960723
      },
      "location":4276,
      "nameEN":"Esgair Carnau",
      "nameCY":"Esgair Carnau",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4276",
      "ngr":"SN9695413333",
      "titleEn":"River Hepste at Esgair Carnau",
      "titleCy":"Afon Hepste yn Esgair Carnau",
      "parameters":[
         {
            "parameter":10083,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.403,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.416653,
         "longitude":-3.4780151
      },
      "location":4277,
      "nameEN":"Frampton Court",
      "nameCY":"Frampton Court",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4277",
      "ngr":"SS9731669663",
      "titleEn":"Llanmaes Brook at Frampton Court",
      "titleCy":"Nant Llanmaes yn Frampton Court",
      "parameters":[
         {
            "parameter":10084,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.544,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.664085,
         "longitude":-3.9056726
      },
      "location":4278,
      "nameEN":"Clarion Close",
      "nameCY":"Clarion Close",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4278",
      "ngr":"SS6829697865",
      "titleEn":"Nant y Fendrod at Clarion Close",
      "titleCy":"Nant y Fendrod yn Clarion Close",
      "parameters":[
         {
            "parameter":10085,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.556,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.685948,
         "longitude":-4.1920149
      },
      "location":4279,
      "nameEN":"Bassett Terrace",
      "nameCY":"Teras Bassett",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4279",
      "ngr":"SN4856700852",
      "titleEn":"Dulais at Bassett Terrace",
      "titleCy":"Afon Dulais yn Teras Bassett",
      "parameters":[
         {
            "parameter":10086,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.738,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.484462,
         "longitude":-3.9826378
      },
      "location":4280,
      "nameEN":"Tal y Bont Ceulan",
      "nameCY":"Tal y Bont Ceulan",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4280",
      "ngr":"SN6546589242",
      "titleEn":"Ceulan at Tal y Bont",
      "titleCy":"Ceulan yn Nhal y Bont",
      "parameters":[
         {
            "parameter":10087,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.320,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.947783,
         "longitude":-3.6893316
      },
      "location":4281,
      "nameEN":"Usk Reservoir",
      "nameCY":"Cronfa Ddwr Wysg",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4281",
      "ngr":"SN8398929047",
      "titleEn":"Usk at Usk Reservoir",
      "titleCy":"Afon Wysg wrth Gronfa Ddwr Wysg",
      "parameters":[
         {
            "parameter":10088,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.161,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.790418,
         "longitude":-3.434077
      },
      "location":4282,
      "nameEN":"Llwynon Compensation Flume",
      "nameCY":"Sianel Gydadferol Llwynon",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4282",
      "ngr":"SO0118711169",
      "titleEn":"Llwynon Compensation Flume",
      "titleCy":"Sianel Gydadferol Llwynon",
      "parameters":[
         {
            "parameter":10089,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":1.297,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.796014,
         "longitude":-3.3635462
      },
      "location":4283,
      "nameEN":"Pontsticill ",
      "nameCY":"Pontsticill ",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4283",
      "ngr":"SO0606311698",
      "titleEn":"Pontsticill",
      "titleCy":"Pontsticill",
      "parameters":[
         {
            "parameter":10090,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.199,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.51308,
         "longitude":-3.5717394
      },
      "location":4284,
      "nameEN":"Springfield Gardens",
      "nameCY":"Gerddi Springfield",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4284",
      "ngr":"SS9102980522",
      "titleEn":"Morfa Brook at Springfield Gardens",
      "titleCy":"Nant y Morfa yn Gerddi Springfield",
      "parameters":[
         {
            "parameter":10091,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":0.500,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.585289,
         "longitude":-3.7746105
      },
      "location":4285,
      "nameEN":"Westend",
      "nameCY":"Westend",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4285",
      "ngr":"SS7714788874",
      "titleEn":"Ffrwdwyllt at Westend",
      "titleCy":"Ffrwdwyllt yn Westend               ",
      "parameters":[
         {
            "parameter":10092,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.657,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.948274,
         "longitude":-3.3951515
      },
      "location":4286,
      "nameEN":"Brecon Promenade",
      "nameCY":"Promenâd Aberhonddu",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4286",
      "ngr":"SO0420828673",
      "titleEn":"River Usk at Brecon Promenade",
      "titleCy":"Afon Wysg wrth Bromenâd Aberhonddu",
      "parameters":[
         {
            "parameter":10093,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.886,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.518604,
         "longitude":-4.0401846
      },
      "location":4287,
      "nameEN":"Ynys Las",
      "nameCY":"Ynys Las",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4287",
      "ngr":"SN6166593148",
      "titleEn":"Afon Leri at Ynys Las",
      "titleCy":"Afon Leri yn Ynys Las",
      "parameters":[
         
      ]
   },
   {
      "coordinates":{
         "latitude":51.804122,
         "longitude":-4.9591693
      },
      "location":4288,
      "nameEN":"Cartlett Brook",
      "nameCY":"Cartlett Brook",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4288",
      "ngr":"SM9607815863",
      "titleEn":"Cartlett Brook at Cartlett (Haverfordwest)",
      "titleCy":"Cartlett Brook yn Cartlett (Hwlffordd)",
      "parameters":[
         {
            "parameter":10094,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":0.514,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":0.0,
         "longitude":0.0
      },
      "location":4289,
      "nameEN":"Clywedog Reservoir Level",
      "nameCY":"Lefel Llyn Clywedog",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4289",
      "ngr":"SN9125287076",
      "titleEn":"Clywedog Reservoir Level",
      "titleCy":"Lefel Llyn Clywedog",
      "parameters":[
         {
            "parameter":10255,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Steady",
            "parameterStatusCY":"Lefel Sefydlog",
            "units":"m",
            "latestValue":-2.285,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":0.0,
         "longitude":0.0
      },
      "location":4291,
      "nameEN":"Dwyran Braint",
      "nameCY":"Dwyran Afon Braint",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4291",
      "ngr":"SH4431565546",
      "titleEn":"River Braint at Dwyran",
      "titleCy":"Afon Braint ar Dwyran",
      "parameters":[
         {
            "parameter":10256,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":2.174,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":0.0,
         "longitude":0.0
      },
      "location":4292,
      "nameEN":"Llangefni",
      "nameCY":"Llangefni",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4292",
      "ngr":"SH4548875758",
      "titleEn":"Cefni at Llangefni",
      "titleCy":"Afon Cefni at Llangefni",
      "parameters":[
         {
            "parameter":10257,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":8.554,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":0.0,
         "longitude":0.0
      },
      "location":4293,
      "nameEN":"Cwm, Elm Street",
      "nameCY":"Cwm,Stryd Elm",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4293",
      "ngr":"SO1805905691",
      "titleEn":"River Ebbw at Cwm, Elm Street",
      "titleCy":"Afon Ebwy yn Cwm, Stryd Elm",
      "parameters":[
         {
            "parameter":10258,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":0.352,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":0.0,
         "longitude":0.0
      },
      "location":4294,
      "nameEN":"Dwyran Rhyd Y Valley",
      "nameCY":"Dwyran Rhyd-Y-Valley",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4294",
      "ngr":"SH4538165669",
      "titleEn":"Afon Rhyd Y Valley at Dwyran",
      "titleCy":"Afon Rhyd-Y-Valley yn Nwyran",
      "parameters":[
         {
            "parameter":10259,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"m",
            "latestValue":2.480,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":0.0,
         "longitude":0.0
      },
      "location":4295,
      "nameEN":"Bowling Bank",
      "nameCY":"Bowling Bank",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"4295",
      "ngr":"SJ3958048249",
      "titleEn":"River Clywedog at Bowling Bank",
      "titleCy":"Afon Clwyedog yn Bowling Bank",
      "parameters":[
         {
            "parameter":10260,
            "paramNameEN":"River Level",
            "paramNameCY":"Lefel yr afon",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"m",
            "latestValue":1.182,
            "latestTime":"2020-10-06T20:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.404931,
         "longitude":-3.0135621
      },
      "location":70139,
      "nameEN":"Liverpool",
      "nameCY":"Liverpool",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"70139",
      "ngr":"SJ3271490269",
      "titleEn":"Liverpool tide gauge",
      "titleCy":"Mesurydd llanw Lerpwl",
      "parameters":[
         {
            "parameter":10250,
            "paramNameEN":"Tide Level",
            "paramNameCY":"Lefel y Llanw",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"mAOD",
            "latestValue":-1.571,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.310472,
         "longitude":-4.6284417
      },
      "location":70539,
      "nameEN":"UKCMF HOLYHEAD TS",
      "nameCY":"UKCMF HOLYHEAD TS",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"70539",
      "ngr":"SH2498282504",
      "titleEn":"UKCMF HOLYHEAD TS tide gauge",
      "titleCy":"Mesurydd llanw Caergybi",
      "parameters":[
         {
            "parameter":10248,
            "paramNameEN":"Tide Level",
            "paramNameCY":"Lefel y Llanw",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"mAOD",
            "latestValue":0.092,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.210566,
         "longitude":-4.1124861
      },
      "location":72439,
      "nameEN":"Ilfracombe",
      "nameCY":"Ilfracombe",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"72439",
      "ngr":"SS5253947828",
      "titleEn":"Ilfracombe tide gauge",
      "titleCy":"Mesurydd llanw Ilfracombe",
      "parameters":[
         {
            "parameter":10249,
            "paramNameEN":"Tide Level",
            "paramNameCY":"Lefel y Llanw",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"mAOD",
            "latestValue":3.195,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.492687,
         "longitude":-2.7565905
      },
      "location":72639,
      "nameEN":"Avonmouth Portbury",
      "nameCY":"Avonmouth Portbury",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"72639",
      "ngr":"ST4757077354",
      "titleEn":"Avonmouth Portbury tide gauge",
      "titleCy":"Mesurydd llanw Avonmouth Portbury",
      "parameters":[
         {
            "parameter":10245,
            "paramNameEN":"Tide Level",
            "paramNameCY":"Lefel y Llanw",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"mAOD",
            "latestValue":5.441,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.453189,
         "longitude":-3.159917
      },
      "location":72839,
      "nameEN":"Newport",
      "nameCY":"Newport",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"72839",
      "ngr":"ST3162983920",
      "titleEn":"Newport tide gauge",
      "titleCy":"Mesurydd llanw Caesnewydd",
      "parameters":[
         {
            "parameter":10254,
            "paramNameEN":"Tide Level",
            "paramNameCY":"Lefel y Llanw",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"mAOD",
            "latestValue":5.197,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.569605,
         "longitude":-3.9776775
      },
      "location":72939,
      "nameEN":"Mumbles",
      "nameCY":"Mumbles",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"72939",
      "ngr":"SS6303287491",
      "titleEn":"Mumbles tide gauge",
      "titleCy":"Mesurydd llanw Mumbles",
      "parameters":[
         {
            "parameter":10253,
            "paramNameEN":"Tide Level",
            "paramNameCY":"Lefel y Llanw",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"mAOD",
            "latestValue":3.511,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":51.707031,
         "longitude":-5.0523641
      },
      "location":73039,
      "nameEN":"Milford Haven",
      "nameCY":"Milford Haven",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"73039",
      "ngr":"SM8920205333",
      "titleEn":"Milford Haven tide gauge",
      "titleCy":"Mesurydd llanw Aberdaugleddau",
      "parameters":[
         {
            "parameter":10252,
            "paramNameEN":"Tide Level",
            "paramNameCY":"Lefel y Llanw",
            "parameterStatusEN":"Level Falling",
            "parameterStatusCY":"Lefel cwympo",
            "units":"mAOD",
            "latestValue":2.535,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.006943,
         "longitude":-4.9899699
      },
      "location":73139,
      "nameEN":"Fishguard",
      "nameCY":"Fishguard",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"73139",
      "ngr":"SM9488238501",
      "titleEn":"Fishguard tide gauge",
      "titleCy":"Mesurydd llanw Abergwaun",
      "parameters":[
         {
            "parameter":10247,
            "paramNameEN":"Tide Level",
            "paramNameCY":"Lefel y Llanw",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"mAOD",
            "latestValue":2.033,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":52.719266,
         "longitude":-4.0544707
      },
      "location":73239,
      "nameEN":"Barmouth",
      "nameCY":"Barmouth",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"73239",
      "ngr":"SH6133215493",
      "titleEn":"Barmouth tide gauge",
      "titleCy":"Mesurydd llanw Abermaw",
      "parameters":[
         {
            "parameter":10246,
            "paramNameEN":"Tide Level",
            "paramNameCY":"Lefel y Llanw",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"mAOD",
            "latestValue":1.734,
            "latestTime":"2020-10-06T21:15:00Z"
         }
      ]
   },
   {
      "coordinates":{
         "latitude":53.327233,
         "longitude":-3.8300191
      },
      "location":73339,
      "nameEN":"Llandudno",
      "nameCY":"Llandudno",
      "statusEN":"Online",
      "statusCY":"Ar-lein",
      "url":"73339",
      "ngr":"SH7821782708",
      "titleEn":"Llandudno tide gauge",
      "titleCy":"Mesurydd llanw Llandudno",
      "parameters":[
         {
            "parameter":10251,
            "paramNameEN":"Tide Level",
            "paramNameCY":"Lefel y Llanw",
            "parameterStatusEN":"Level Rising",
            "parameterStatusCY":"Lefel codi",
            "units":"mAOD",
            "latestValue":-1.359,
            "latestTime":"2020-10-06T21:00:00Z"
         }
      ]
   }
]`
