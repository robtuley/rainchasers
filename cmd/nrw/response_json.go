package main

const jsonResponseFromApi = `
{
  "type" : "FeatureCollection", 
  "crs" : 
  {
    "type" : "name", 
    "properties" : 
    {
      "name" : "EPSG:4326"
    }
  }, 
  "features" : [
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.16663310593962, 51.6973802584378
        ]
      }, 
      "properties" : {
        "Location" : "4268", 
        "LatestValue" : "0.214", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Trebeddrod", 
        "NameCY" : "Trebeddrod", 
        "NGR" : "SN5036002070", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Offline", 
        "StatusCY" : "All-lein", 
        "TitleEN" : "Trebeddrod Reservoir Level", 
        "TitleCY" : "Trebeddrod Cronfa Ddwr Lefel", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4268"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.08019174558596, 53.1006247513997
        ]
      }, 
      "properties" : {
        "Location" : "4177", 
        "LatestValue" : "0.873", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Nant Peris", 
        "NameCY" : "Nant Peris", 
        "NGR" : "SH6082557958", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Offline", 
        "StatusCY" : "All-lein", 
        "TitleEN" : "Peris at Nant Peris", 
        "TitleCY" : "Afon Peris yn Nant Peris", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4177"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.67056558678914, 51.6464877290809
        ]
      }, 
      "properties" : {
        "Location" : "4009", 
        "LatestValue" : "-5.451", 
        "LatestTime" : "24/10/2017 08:30", 
        "NameEN" : "Chepstow Bridge", 
        "NameCY" : "Pont Cas-gwent", 
        "NGR" : "ST5370094400", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Offline", 
        "StatusCY" : "All-lein", 
        "TitleEN" : "Wye at Chepstow Bridge", 
        "TitleCY" : "Afon Gwy ym Mhont Cas-gwent", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4009"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.44572425903852, 52.0462808165848
        ]
      }, 
      "properties" : {
        "Location" : "4052", 
        "LatestValue" : "0.182", 
        "LatestTime" : "25/09/2018 08:00", 
        "NameEN" : "Upper Chapel", 
        "NameCY" : "Capel Uchaf", 
        "NGR" : "SO0095039640", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Offline", 
        "StatusCY" : "All-lein", 
        "TitleEN" : "Honddu at Upper Chapel", 
        "TitleCY" : "Afon Honddu yng Nghapel Uchaf", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4052"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.87137949358349, 53.1341524390619
        ]
      }, 
      "properties" : {
        "Location" : "4173", 
        "LatestValue" : "5.798", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Ironbridge", 
        "NameCY" : "Ironbridge", 
        "NGR" : "SJ4180060020", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dee at Ironbridge", 
        "TitleCY" : "Afon Ddyfrdwy yn Ironbridge", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4173"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.8470514914182, 53.1200261615446
        ]
      }, 
      "properties" : {
        "Location" : "4174", 
        "LatestValue" : "1.524", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Lea Hall", 
        "NameCY" : "Lea Hall", 
        "NGR" : "SJ4340958429", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Aldford brook at Lea Hall", 
        "TitleCY" : "Aldford Brook yn Lea Hall", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4174"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.79722869286302, 53.1371147919168
        ]
      }, 
      "properties" : {
        "Location" : "4175", 
        "LatestValue" : "5.188", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Pont Fawr", 
        "NameCY" : "Pont Fawr", 
        "NGR" : "SH7987061504", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Conwy at Pont Fawr", 
        "TitleCY" : "Afon Conwy ym Mhont Fawr", 
        "Units" : "mAOD", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4175"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.23671845639978, 52.9624521227109
        ]
      }, 
      "properties" : {
        "Location" : "4171", 
        "LatestValue" : "0.423", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Garndolbenmaen", 
        "NameCY" : "Garndolbenmaen", 
        "NGR" : "SH4986742906", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dwyfor at Garndolbenmaen", 
        "TitleCY" : "Afon Dwyfor yng Ngarndolbenmaen", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4171"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.97163670304611, 52.9666580214283
        ]
      }, 
      "properties" : {
        "Location" : "4176", 
        "LatestValue" : "1.422", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Manley Hall", 
        "NameCY" : "Manley Hall", 
        "NGR" : "SJ3484041474", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dee at Manley Hall", 
        "TitleCY" : "Afon Ddyfrdwy ym Manley Hall", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4176"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.87914121827605, 53.0830991035745
        ]
      }, 
      "properties" : {
        "Location" : "4170", 
        "LatestValue" : "8.092", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Farndon", 
        "NameCY" : "Farndon", 
        "NGR" : "SJ4121154347", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dee at Farndon", 
        "TitleCY" : "Afon Ddyfrdwy yn Rhedynfre", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4170"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.67459548716969, 52.8629934405363
        ]
      }, 
      "properties" : {
        "Location" : "4178", 
        "LatestValue" : "0.913", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "New Inn", 
        "NameCY" : "New Inn", 
        "NGR" : "SH8736130814", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dee at New Inn", 
        "TitleCY" : "Afon Ddyfrdwy yn New Inn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4178"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.25204147199603, 53.1363814763128
        ]
      }, 
      "properties" : {
        "Location" : "4179", 
        "LatestValue" : "0.890", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Peblic Mill", 
        "NameCY" : "Melin Peblig", 
        "NGR" : "SH4944662283", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Seiont at Peblic Mill", 
        "TitleCY" : "Afon Seiont ym Melin Peblig", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4179"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.43447503675377, 53.2624644169437
        ]
      }, 
      "properties" : {
        "Location" : "4180", 
        "LatestValue" : "2.879", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Pont Dafydd", 
        "NameCY" : "Pont Dafydd", 
        "NGR" : "SJ0441574900", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Clwyd at Pont Dafydd", 
        "TitleCY" : "Afon Clwyd ym Mhont Dafydd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4180"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.39476551389685, 53.2276427536778
        ]
      }, 
      "properties" : {
        "Location" : "4181", 
        "LatestValue" : "1.521", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Pont y Cambwll", 
        "NameCY" : "Pont y Cambwll", 
        "NGR" : "SJ0698870974", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Clwyd at Pont y Cambwll", 
        "TitleCY" : "Afon Clwyd ym Mhont y Cambwll", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4181"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.99374556617948, 53.0794102975451
        ]
      }, 
      "properties" : {
        "Location" : "4182", 
        "LatestValue" : "0.537", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Pont y Capel", 
        "NameCY" : "Pont y Capel", 
        "NGR" : "SJ3352954037", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Alyn at Pont y Capel", 
        "TitleCY" : "Afon Alun ym Mhont y Capel", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4182"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.57095988431466, 53.2327225832398
        ]
      }, 
      "properties" : {
        "Location" : "4183", 
        "LatestValue" : "0.877", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Pont y Gwyddel", 
        "NameCY" : "Pont y Gwyddel", 
        "NGR" : "SH9523971783", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Elwy at Pont y Gwyddel", 
        "TitleCY" : "Afon Elwy ym Mhont y Gwyddel", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4183"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.47501956126953, 53.293496047503
        ]
      }, 
      "properties" : {
        "Location" : "4185", 
        "LatestValue" : "2.980", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Rhuddlan", 
        "NameCY" : "Rhuddlan", 
        "NGR" : "SJ0178278407", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Clwyd at Rhuddlan", 
        "TitleCY" : "Afon Clwyd yn Rhuddlan", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4185"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.95273409308966, 51.8982777105896
        ]
      }, 
      "properties" : {
        "Location" : "4205", 
        "LatestValue" : "2.208", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Manorafon", 
        "NameCY" : "Manorafon", 
        "NGR" : "SN6574023994", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Tywi at Manorafon", 
        "TitleCY" : "Afon Tywi ym Manorafon", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4205"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.76444162097235, 52.1217832690926
        ]
      }, 
      "properties" : {
        "Location" : "4204", 
        "LatestValue" : "0.164", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Llyn Brianne Reservoir", 
        "NameCY" : "Argae Llyn Brianne", 
        "NGR" : "SN7929848520", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Tywi at Llyn Brianne Reservoir", 
        "TitleCY" : "Afon Tywi yn Argae Llyn Brianne", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4204"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.87258297801958, 51.9673262523653
        ]
      }, 
      "properties" : {
        "Location" : "4203", 
        "LatestValue" : "0.577", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Llanwrda", 
        "NameCY" : "Llanwrda", 
        "NGR" : "SN7145231528", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dulais at Llanwrda", 
        "TitleCY" : "Afon Dulais yn Llanwrda", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4203"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.88438637782952, 51.9401980511352
        ]
      }, 
      "properties" : {
        "Location" : "4202", 
        "LatestValue" : "0.773", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Llangadog", 
        "NameCY" : "Llangadog", 
        "NGR" : "SN7056328532", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Bran at Llangadog", 
        "TitleCY" : "Afon Brân yn Llangadog", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4202"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.79100080596327, 51.9938897536961
        ]
      }, 
      "properties" : {
        "Location" : "4201", 
        "LatestValue" : "0.720", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Llandovery", 
        "NameCY" : "Llanymddyfri", 
        "NGR" : "SN7712934341", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Bran at Llandovery", 
        "TitleCY" : "Afon Brân yn Llanymddyfri", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4201"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.28151642131484, 51.8739519232382
        ]
      }, 
      "properties" : {
        "Location" : "4199", 
        "LatestValue" : "1.057", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Glan Gwili", 
        "NameCY" : "Glan Gwili", 
        "NGR" : "SN4303621947", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Gwili at Glan Gwili", 
        "TitleCY" : "Afon Gwili yng Nglan Gwili", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4199"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.87421675573465, 51.9227452353643
        ]
      }, 
      "properties" : {
        "Location" : "4198", 
        "LatestValue" : "0.520", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Felin Y Cwm", 
        "NameCY" : "Felin y Cwm", 
        "NGR" : "SN7121226573", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Sawdde at Felin Y Cwm", 
        "TitleCY" : "Afon Sawdde yn Felin y Cwm", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4198"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.16804286396231, 51.8809816163193
        ]
      }, 
      "properties" : {
        "Location" : "4197", 
        "LatestValue" : "1.492", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Felin Mynachdy", 
        "NameCY" : "Felin Mynachdy", 
        "NGR" : "SN5087022490", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Cothi at Felin Mynachdy", 
        "TitleCY" : "Afon Cothi yn Felin Mynachdy", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4197"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.80540161712309, 52.0107811360719
        ]
      }, 
      "properties" : {
        "Location" : "4196", 
        "LatestValue" : "0.907", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Dolau Hirion", 
        "NameCY" : "Dolau Hirion", 
        "NGR" : "SN7618736244", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Tywi at Dolau Hirion", 
        "TitleCY" : "Afon Tywi yn Nolau Hirion", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4196"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.9710973519009, 52.0777753499051
        ]
      }, 
      "properties" : {
        "Location" : "4195", 
        "LatestValue" : "0.564", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Ddol Las", 
        "NameCY" : "Ddol Las", 
        "NGR" : "SN6501843990", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Twrch at Ddol Las", 
        "TitleCY" : "Afon Twrch yn Nôl-Las", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4195"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.79247385631598, 52.1094194630717
        ]
      }, 
      "properties" : {
        "Location" : "4194", 
        "LatestValue" : "0.728", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Craig Clungwyn", 
        "NameCY" : "Craig Clungwyn", 
        "NGR" : "SN7734547192", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Doethie at Craig Clungwyn", 
        "TitleCY" : "Afon Doethie yng Nghraig Clungwyn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4194"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.20147699475324, 51.8631489147549
        ]
      }, 
      "properties" : {
        "Location" : "4193", 
        "LatestValue" : "2.631", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Capel Dewi", 
        "NameCY" : "Capel Dewi", 
        "NGR" : "SN4850920576", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Tywi at Capel Dewi", 
        "TitleCY" : "Afon Tywi yng Nghapel Dewi", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4193"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.86629875952559, 53.0078353513566
        ]
      }, 
      "properties" : {
        "Location" : "4192", 
        "LatestValue" : "1.179", 
        "LatestTime" : "13/03/2019 19:30", 
        "NameEN" : "Wych Brook", 
        "NameCY" : "Nant Wych", 
        "NGR" : "SJ4197045964", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Wych Brook at Wych Brook", 
        "TitleCY" : "Nant Wych yn Nant Wych", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4192"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.82413088005124, 53.1602149531495
        ]
      }, 
      "properties" : {
        "Location" : "4191", 
        "LatestValue" : "2.372", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Trefriw", 
        "NameCY" : "Trefriw", 
        "NGR" : "SH7813664119", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Conwy at Trefriw", 
        "TitleCY" : "Afon Conwy yn Nhrefriw", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4191"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.44956083695488, 53.2621110999647
        ]
      }, 
      "properties" : {
        "Location" : "4188", 
        "LatestValue" : "1.654", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "St Asaph", 
        "NameCY" : "Llanelwy", 
        "NGR" : "SJ0340874881", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Elwy at St Asaph", 
        "TitleCY" : "Afon Elwy yn Llanelwy", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4188"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.31375140491492, 53.1228822945446
        ]
      }, 
      "properties" : {
        "Location" : "4187", 
        "LatestValue" : "0.627", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Ruthin Weir", 
        "NameCY" : "Cored Rhuthun", 
        "NGR" : "SJ1218259218", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Clwyd at Ruthin Weir", 
        "TitleCY" : "Afon Clwyd yng Nghored Rhuthun", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4187"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.19280913825219, 53.1918555042545
        ]
      }, 
      "properties" : {
        "Location" : "4186", 
        "LatestValue" : "0.954", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Rhydymwyn", 
        "NameCY" : "Rhydymwyn", 
        "NGR" : "SJ2040366749", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Alyn at Rhydymwyn", 
        "TitleCY" : "Afon Alun yn Rhyd-y-mwyn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4186"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.97005830034082, 51.820376651159
        ]
      }, 
      "properties" : {
        "Location" : "4124", 
        "LatestValue" : "1.339", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Prendergast Mill", 
        "NameCY" : "Melin Prendergast", 
        "NGR" : "SM9540217700", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Western Cleddau at Prendergast Mill", 
        "TitleCY" : "Afon Cleddau Wen ym Melin Prendergast", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4124"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.71575233304566, 51.70148849291
        ]
      }, 
      "properties" : {
        "Location" : "4125", 
        "LatestValue" : "0.791", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Resolven", 
        "NameCY" : "Resolfen", 
        "NGR" : "SN8152901698", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Neath at Resolven", 
        "TitleCY" : "Afon Nedd yn Resolfen", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4125"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.49630724475289, 51.8199919518917
        ]
      }, 
      "properties" : {
        "Location" : "4126", 
        "LatestValue" : "0.804", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "St Clears", 
        "NameCY" : "Sanclêr", 
        "NGR" : "SN2804616432", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Cynin at St Clears", 
        "TitleCY" : "Afon Cynin yn Sanclêr", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4126"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.69841024489968, 51.7367161685326
        ]
      }, 
      "properties" : {
        "Location" : "4127", 
        "LatestValue" : "0.329", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Stepaside", 
        "NameCY" : "Stepaside", 
        "NGR" : "SN1377607669", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Fords Lake at Stepaside", 
        "TitleCY" : "Fords Lake yn Stepaside", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4127"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.13063664313552, 52.1846377826151
        ]
      }, 
      "properties" : {
        "Location" : "4129", 
        "LatestValue" : "0.679", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Talsarn", 
        "NameCY" : "Tal-sarn", 
        "NGR" : "SN5443556183", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Aeron at Talsarn", 
        "TitleCY" : "Afon Aeron yn Nhal-sarn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4129"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.99775099596336, 51.7957118581956
        ]
      }, 
      "properties" : {
        "Location" : "4131", 
        "LatestValue" : "0.940", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Tir Y Dail", 
        "NameCY" : "Tir y Dail", 
        "NGR" : "SN6233012672", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Loughor at Tir Y Dail", 
        "TitleCY" : "Afon Llwchwr yn Nhir y Dail", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4131"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.60961303975858, 51.8190054376294
        ]
      }, 
      "properties" : {
        "Location" : "4132", 
        "LatestValue" : "0.531", 
        "LatestTime" : "13/03/2019 19:15", 
        "NameEN" : "Whitland", 
        "NameCY" : "Hendy-Gwyn", 
        "NGR" : "SN2023416596", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Gronw at Whitland", 
        "TitleCY" : "Afon Gronw yn Hendy-gwyn ar Daf", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4132"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.69564786943242, 51.7122292595859
        ]
      }, 
      "properties" : {
        "Location" : "4133", 
        "LatestValue" : "0.299", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Ynys Fach Bridge", 
        "NameCY" : "Pont Ynys Fach", 
        "NGR" : "SN8294602860", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Clydach Brook at Ynys Fach Bridge", 
        "TitleCY" : "Nant Clydach ym Mhont Ynys Fach", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4133"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.90269962518433, 51.6815229453174
        ]
      }, 
      "properties" : {
        "Location" : "4134", 
        "LatestValue" : "0.973", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Ynystanglws", 
        "NameCY" : "Ynystanglws", 
        "NGR" : "SS6855399798", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Tawe at Ynystanglws", 
        "TitleCY" : "Afon Tawe yn Ynystanglws", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4134"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.75001400070845, 52.731460880215
        ]
      }, 
      "properties" : {
        "Location" : "4143", 
        "LatestValue" : "0.227", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Llawr Cae", 
        "NameCY" : "Llawr Cae", 
        "NGR" : "SH8192816305", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Cerist at Llawr Cae", 
        "TitleCY" : "Afon Cerist yn Llawr Cae", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4143"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.12662359497553, 53.0313659932012
        ]
      }, 
      "properties" : {
        "Location" : "4144", 
        "LatestValue" : "0.144", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Hafod Wydr", 
        "NameCY" : "Hafod Wydr", 
        "NGR" : "SH5748850346", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Colwyn at Hafod Wydr", 
        "TitleCY" : "Afon Colwyn yn Hafod Wydr", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4144"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.79220469420364, 53.1054910739533
        ]
      }, 
      "properties" : {
        "Location" : "4145", 
        "LatestValue" : "1.293", 
        "LatestTime" : "13/03/2019 19:15", 
        "NameEN" : "Cwmlanerch", 
        "NameCY" : "Cwmlanerch", 
        "NGR" : "SH8011857978", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Conwy at Cwmlanerch", 
        "TitleCY" : "Afon Conwy yng Nghwm Llannerch", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4145"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.85509738453768, 52.600609536898
        ]
      }, 
      "properties" : {
        "Location" : "4146", 
        "LatestValue" : "1.968", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Dyfi Bridge", 
        "NameCY" : "Pont Dyfi", 
        "NGR" : "SH7445801929", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dyfi at Dyfi Bridge", 
        "TitleCY" : "Afon Dyfi ym Mhont Dyfi", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4146"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.01175961246602, 52.4732965698391
        ]
      }, 
      "properties" : {
        "Location" : "4147", 
        "LatestValue" : "0.343", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Dolybont", 
        "NameCY" : "Dol y bont", 
        "NGR" : "SN6345488054", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Leri at Dolybont", 
        "TitleCY" : "Afon Tyleri yn Nôl-y-bont", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4147"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.81315674321152, 53.0676751714313
        ]
      }, 
      "properties" : {
        "Location" : "4148", 
        "LatestValue" : "0.693", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Pont Gethin", 
        "NameCY" : "Pont Gethin", 
        "NGR" : "SH7860953807", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Lledr at Pont Gethin", 
        "TitleCY" : "Afon Lledr ym Mhont Gethin", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4148"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.32830560842612, 53.0456772228517
        ]
      }, 
      "properties" : {
        "Location" : "4149", 
        "LatestValue" : "0.495", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Pont y Cim", 
        "NameCY" : "Pont y Cim", 
        "NGR" : "SH4401752358", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Llyfni at Pont y Cim", 
        "TitleCY" : "Afon Llyfni ym Mhont y Cim", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4149"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.38216393609564, 52.9368583688443
        ]
      }, 
      "properties" : {
        "Location" : "4169", 
        "LatestValue" : "0.209", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Pencaenewydd", 
        "NameCY" : "Pencaenewydd", 
        "NGR" : "SH4000540374", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Erch at Pencaenewydd", 
        "TitleCY" : "Afon Erch ym Mhencaenewydd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4169"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.86643125136116, 53.0077534808529
        ]
      }, 
      "properties" : {
        "Location" : "4168", 
        "LatestValue" : "1.172", 
        "LatestTime" : "13/03/2019 19:30", 
        "NameEN" : "Emral Brook", 
        "NameCY" : "Emral Brook", 
        "NGR" : "SJ4196145955", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Emral Brook at Emral Brook", 
        "TitleCY" : "Emral Brook yn Emral Brook", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4168"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.0242246091118, 52.6397465863591
        ]
      }, 
      "properties" : {
        "Location" : "4167", 
        "LatestValue" : "0.923", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Pont y Garth", 
        "NameCY" : "Pont y Garth", 
        "NGR" : "SH6312706590", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dysynni at Pont y Garth", 
        "TitleCY" : "Afon Dysynni ym Mhont y Garth", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4167"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.43015973105281, 52.9812626324417
        ]
      }, 
      "properties" : {
        "Location" : "4166", 
        "LatestValue" : "0.881", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Druid", 
        "NameCY" : "Druid", 
        "NGR" : "SJ0407843614", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Alwen at Druid", 
        "TitleCY" : "Afon Alwen yn y Ddwyryd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4166"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.59271354592463, 52.9031022220421
        ]
      }, 
      "properties" : {
        "Location" : "4165", 
        "LatestValue" : "1.878", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Dee Bridge", 
        "NameCY" : "Pont y Dee", 
        "NGR" : "SH9297235150", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dee at Dee Bridge", 
        "TitleCY" : "Afon Ddyfrdwy ym Mhont Dyfrdwy", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4165"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.72491028642943, 52.9632801474328
        ]
      }, 
      "properties" : {
        "Location" : "4164", 
        "LatestValue" : "0.375", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Cynefail", 
        "NameCY" : "Cynefail", 
        "NGR" : "SH8424242049", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "River Gelyn Level at Cynefail", 
        "TitleCY" : "Lefel Afon Gelyn yng Nghynefail", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4164"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.38793885472605, 52.9791561272538
        ]
      }, 
      "properties" : {
        "Location" : "4163", 
        "LatestValue" : "1.139", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Corwen", 
        "NameCY" : "Corwen", 
        "NGR" : "SJ0690843324", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dee at Corwen", 
        "TitleCY" : "Afon Ddyfrdwy yng Nghorwen", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4163"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.0500785359914, 52.9287196192119
        ]
      }, 
      "properties" : {
        "Location" : "4162", 
        "LatestValue" : "0.716", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Brynkinalt", 
        "NameCY" : "Brynkinalt", 
        "NGR" : "SJ2951037328", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ceiriog at Brynkinalt", 
        "TitleCY" : "Afon Ceiriog ym Mryncunallt", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4162"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.81352480908163, 53.0814534140253
        ]
      }, 
      "properties" : {
        "Location" : "4160", 
        "LatestValue" : "0.349", 
        "LatestTime" : "13/03/2019 15:00", 
        "NameEN" : "Carden Brook", 
        "NameCY" : "Carden Brook", 
        "NGR" : "SJ4560454112", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Carden Brook at Carden Park", 
        "TitleCY" : "Carden Brook yn Parc Carden", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4160"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.26692214694896, 53.1146106028814
        ]
      }, 
      "properties" : {
        "Location" : "4158", 
        "LatestValue" : "0.624", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Bontnewydd", 
        "NameCY" : "Bontnewydd", 
        "NGR" : "SH4837459893", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Gwyrfai at Bontnewydd", 
        "TitleCY" : "Afon Gwyrfai yn y Bontnewydd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4158"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.3565432000415, 53.2655996814072
        ]
      }, 
      "properties" : {
        "Location" : "4157", 
        "LatestValue" : "0.546", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Bodffordd", 
        "NameCY" : "Bodfford", 
        "NGR" : "SH4293076880", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Frogwy at Bodffordd", 
        "TitleCY" : "Afon Frogwy ym Modffordd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4157"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.34209214586927, 53.2314120193653
        ]
      }, 
      "properties" : {
        "Location" : "4156", 
        "LatestValue" : "0.348", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Bodfari", 
        "NameCY" : "Bodfari", 
        "NGR" : "SJ1051271326", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Wheeler at Bodfari", 
        "TitleCY" : "Afon Chwiler ym Modfari", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4156"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.10025726942165, 53.0086923895686
        ]
      }, 
      "properties" : {
        "Location" : "4155", 
        "LatestValue" : "0.723", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Beddgelert", 
        "NameCY" : "Beddgelert", 
        "NGR" : "SH5918247772", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Glaslyn at Beddgelert", 
        "TitleCY" : "Afon Glaslyn ym Meddgelert", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4155"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.5897852221493, 52.9104325519676
        ]
      }, 
      "properties" : {
        "Location" : "4153", 
        "LatestValue" : "1.045", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Bala Weir X", 
        "NameCY" : "Gored X Y Bala", 
        "NGR" : "SH9318735961", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Tryweryn at Bala Weir X", 
        "TitleCY" : "Afon Tryweryn yng Nghored X y Bala", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4153"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.57621432485281, 52.9085636740457
        ]
      }, 
      "properties" : {
        "Location" : "4152", 
        "LatestValue" : "1.079", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Bala", 
        "NameCY" : "Y Bala", 
        "NGR" : "SH9409535733", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dee at Bala", 
        "TitleCY" : "Afon Ddyfrdwy yn y Bala", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4152"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.8841350570093, 52.7444107261397
        ]
      }, 
      "properties" : {
        "Location" : "4151", 
        "LatestValue" : "1.881", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Dolgellau", 
        "NameCY" : "Dolgellau", 
        "NGR" : "SH7291017974", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Wnion at Dolgellau", 
        "TitleCY" : "Afon Wnion yn Nolgellau", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4151"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.87709834229229, 52.8201978092437
        ]
      }, 
      "properties" : {
        "Location" : "4150", 
        "LatestValue" : "0.430", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Tyddyn Gwladys", 
        "NameCY" : "Tyddyn Gwladys", 
        "NGR" : "SH7360526391", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Mawddach at Tyddyn Gwladys", 
        "TitleCY" : "Afon Mawddach yn Nhyddyn Gwladys", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4150"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.48853348575924, 52.4552098085026
        ]
      }, 
      "properties" : {
        "Location" : "4260", 
        "LatestValue" : "0.901", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Dolwen", 
        "NameCY" : "Dolwen", 
        "NGR" : "SN9894985181", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "River Severn at Dolwen", 
        "TitleCY" : "Afon Hafren yn Dolwen", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4260"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.26459894220742, 52.7636499669963
        ]
      }, 
      "properties" : {
        "Location" : "4261", 
        "LatestValue" : "0.578", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Llanfyllin (Cain)", 
        "NameCY" : "Llanfyllin (Cain)", 
        "NGR" : "SJ1476719199", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Cain at Llanfyllin", 
        "TitleCY" : "Afon Cain yn Llanfyllin", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4261"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.17421357968211, 52.7759427278334
        ]
      }, 
      "properties" : {
        "Location" : "4262", 
        "LatestValue" : "1.342", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Llansantffraid", 
        "NameCY" : "Llansantffraid", 
        "NGR" : "SJ2088820463", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Cain at Llansantffraid", 
        "TitleCY" : "Afon Cain yn Llansantffraid", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4262"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.16132298231808, 52.5916730081122
        ]
      }, 
      "properties" : {
        "Location" : "4263", 
        "LatestValue" : "1.377", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Pont-y-Gaer", 
        "NameCY" : "Pont-y-Gaer", 
        "NGR" : "SO2142799952", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Camlad at Pont-y-Gaer", 
        "TitleCY" : "Afon Camlad ym Mhont-y-Gaer", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4263"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.96695796079973, 51.8795890675344
        ]
      }, 
      "properties" : {
        "Location" : "4264", 
        "LatestValue" : "0.667", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Treffgarne", 
        "NameCY" : "Treffgarne", 
        "NGR" : "SM9588424275", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Western Cleddau at Treffgarne", 
        "TitleCY" : "Afon Cleddau orllewin yn Nhrefgarn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4264"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.79097836266495, 51.5930249298056
        ]
      }, 
      "properties" : {
        "Location" : "4265", 
        "LatestValue" : "3.285", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Green Park Weir", 
        "NameCY" : "Cored Green Park", 
        "NGR" : "SS7603589761", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "River Afan at Greenpark Weir", 
        "TitleCY" : "Afon Afan yng Nghored Green Park", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4265"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.95275370196628, 51.8330214487512
        ]
      }, 
      "properties" : {
        "Location" : "4266", 
        "LatestValue" : "0.352", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Forge Bridge", 
        "NameCY" : "Pont Forge", 
        "NGR" : "SN6554416737", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Loughor at Forge Bridge", 
        "TitleCY" : "Llwchwr ym Mhont Forge", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4266"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.98262382118228, 52.4831936405391
        ]
      }, 
      "properties" : {
        "Location" : "4267", 
        "LatestValue" : "0.455", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Tal y Bont Leri", 
        "NameCY" : "Tal y Bont Leri", 
        "NGR" : "SN6546389100", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Leri at Tal y Bont", 
        "TitleCY" : "Leri yn Nhal y Bont", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4267"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.61663025814182, 51.7475487884619
        ]
      }, 
      "properties" : {
        "Location" : "4269", 
        "LatestValue" : "0.927", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Glynneath", 
        "NameCY" : "Glyn-nedd", 
        "NGR" : "SN8849206664", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "River Neath at Glynneath", 
        "TitleCY" : "Afon Nedd yng Nglyn-nedd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4269"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.00514719397918, 51.8208721957131
        ]
      }, 
      "properties" : {
        "Location" : "4270", 
        "LatestValue" : "0.418", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Llandybie", 
        "NameCY" : "Llandybie", 
        "NGR" : "SN6189715484", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "River Marlas at Llandybie", 
        "TitleCY" : "Afon Marlas yn Llandybie", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4270"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.76112760758422, 51.7724207799259
        ]
      }, 
      "properties" : {
        "Location" : "4271", 
        "LatestValue" : "0.621", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Ystradgynlais", 
        "NameCY" : "Ystradgynlais", 
        "NGR" : "SN7858409661", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Tawe at Ystradgynlais", 
        "TitleCY" : "Afon Tawe yn Ystradgynlais", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4271"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.08123318196923, 51.6978728575757
        ]
      }, 
      "properties" : {
        "Location" : "4272", 
        "LatestValue" : "0.510", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Llangennech", 
        "NameCY" : "Llangennech", 
        "NGR" : "SN5626301953", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "River Morlais at Llangennech", 
        "TitleCY" : "Afon Morlais yn Llangennech", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4272"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.18937075870349, 51.6893596785415
        ]
      }, 
      "properties" : {
        "Location" : "4273", 
        "LatestValue" : "0.381", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Stradey Home Farm", 
        "NameCY" : "Fferm garter y Strade ", 
        "NGR" : "SN4876201225", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dulais at Stradey Home Farm", 
        "TitleCY" : "Dulais yn Fferm Gartre'r Strade", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4273"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.53137727526098, 51.7992232727052
        ]
      }, 
      "properties" : {
        "Location" : "4274", 
        "LatestValue" : "0.317", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Llanddowror", 
        "NameCY" : "Llanddowror", 
        "NGR" : "SN2554914206", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Llanddowror Brook at Llanddowror", 
        "TitleCY" : "Nant Llanddowror yn Llanddowror", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4274"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.45980000712406, 51.4105855828611
        ]
      }, 
      "properties" : {
        "Location" : "4275", 
        "LatestValue" : "0.620", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Llanmaes", 
        "NameCY" : "Llanmaes", 
        "NGR" : "SS9857068962", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Llanmaes Brook at Llanmaes", 
        "TitleCY" : "Nant Llanmaes yn Llanmaes", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4275"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.49608513556602, 51.8091115995248
        ]
      }, 
      "properties" : {
        "Location" : "4276", 
        "LatestValue" : "0.272", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Esgair Carnau", 
        "NameCY" : "Esgair Carnau", 
        "NGR" : "SN9695413333", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "River Hepste at Esgair Carnau", 
        "TitleCY" : "Afon Hepste yn Esgair Carnau", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4276"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.47802734355623, 51.4166606344064
        ]
      }, 
      "properties" : {
        "Location" : "4277", 
        "LatestValue" : "0.439", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Frampton Court", 
        "NameCY" : "Frampton Court", 
        "NGR" : "SS9731669663", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Llanmaes Brook at Frampton Court", 
        "TitleCY" : "Nant Llanmaes yn Frampton Court", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4277"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.90568525454533, 51.6640923275258
        ]
      }, 
      "properties" : {
        "Location" : "4278", 
        "LatestValue" : "0.307", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Clarion Close", 
        "NameCY" : "Clarion Close", 
        "NGR" : "SS6829697865", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Nant y Fendrod at Clarion Close", 
        "TitleCY" : "Nant y Fendrod yn Clarion Close", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4278"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.1920276361203, 51.6859558304325
        ]
      }, 
      "properties" : {
        "Location" : "4279", 
        "LatestValue" : "0.754", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Bassett Terrace", 
        "NameCY" : "Teras Bassett", 
        "NGR" : "SN4856700852", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dulais at Bassett Terrace", 
        "TitleCY" : "Afon Dulais yn Teras Bassett", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4279"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.9826517771426, 52.4844699420033
        ]
      }, 
      "properties" : {
        "Location" : "4280", 
        "LatestValue" : "0.338", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Tal y Bont Ceulan", 
        "NameCY" : "Tal y Bont Ceulan", 
        "NGR" : "SN6546589242", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ceulan at Tal y Bont", 
        "TitleCY" : "Ceulan yn Nhal y Bont", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4280"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.68934466772945, 51.9477903303331
        ]
      }, 
      "properties" : {
        "Location" : "4281", 
        "LatestValue" : "0.383", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Usk Reservoir", 
        "NameCY" : "Cronfa Ddŵr Wysg", 
        "NGR" : "SN8398929047", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Usk at Usk Reservoir", 
        "TitleCY" : "Afon Wysg wrth Gronfa Ddwr Wysg", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4281"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.434089766848, 51.7904263072903
        ]
      }, 
      "properties" : {
        "Location" : "4282", 
        "LatestValue" : "1.055", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Llwynon Compensation Flume", 
        "NameCY" : "Sianel Gydadferol Llwynon", 
        "NGR" : "SO0118711169", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Llwynon Compensation Flume", 
        "TitleCY" : "Sianel Gydadferol Llwynon", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4282"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.36355901095099, 51.7960222777109
        ]
      }, 
      "properties" : {
        "Location" : "4283", 
        "LatestValue" : "0.393", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Pontsticill ", 
        "NameCY" : "Pontsticill ", 
        "NGR" : "SO0606311698", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Pontsticill", 
        "TitleCY" : "Pontsticill", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4283"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.57175177339546, 51.5130879418767
        ]
      }, 
      "properties" : {
        "Location" : "4284", 
        "LatestValue" : "0.488", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Springfield Gardens", 
        "NameCY" : "Gerddi Springfield", 
        "NGR" : "SS9102980522", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Morfa Brook at Springfield Gardens", 
        "TitleCY" : "Nant y Morfa yn Gerddi Springfield", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4284"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.77462299269288, 51.585297268211
        ]
      }, 
      "properties" : {
        "Location" : "4285", 
        "LatestValue" : "0.542", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Westend", 
        "NameCY" : "Westend", 
        "NGR" : "SS7714788874", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ffrwdwyllt at Westend", 
        "TitleCY" : "Ffrwdwyllt yn Westend               ", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4285"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.39516451323965, 51.9482824614861
        ]
      }, 
      "properties" : {
        "Location" : "4286", 
        "LatestValue" : "0.999", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Brecon Promenade", 
        "NameCY" : "Promenâd Aberhonddu", 
        "NGR" : "SO0420828673", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "River Usk at Brecon Promenade", 
        "TitleCY" : "Afon Wysg wrth Bromenâd Aberhonddu", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4286"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.95918240710075, 51.8041294443289
        ]
      }, 
      "properties" : {
        "Location" : "4288", 
        "LatestValue" : "0.489", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Cartlett Brook", 
        "NameCY" : "Cartlett Brook", 
        "NGR" : "SM9607815863", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Cartlett Brook at Cartlett (Haverfordwest)", 
        "TitleCY" : "Cartlett Brook yn Cartlett (Hwlffordd)", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4288"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.60237652165803, 52.4707594927638
        ]
      }, 
      "properties" : {
        "Location" : "4289", 
        "LatestValue" : "0.037", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Clywedog Reservoir Level", 
        "NameCY" : "Lefel Llyn Clywedog", 
        "NGR" : "SN9125287076", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Clywedog Reservoir Level", 
        "TitleCY" : "Lefel Llyn Clywedog", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4289"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.33026164661818, 53.1642132336763
        ]
      }, 
      "properties" : {
        "Location" : "4291", 
        "LatestValue" : "1.761", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Dwyran Braint", 
        "NameCY" : "Dwyran Afon Braint", 
        "NGR" : "SH4431565546", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "River Braint at Dwyran", 
        "TitleCY" : "Afon Braint ar Dwyran", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4291"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.31767863055247, 53.2562741767195
        ]
      }, 
      "properties" : {
        "Location" : "4292", 
        "LatestValue" : "8.775", 
        "LatestTime" : "12/03/2019 10:30", 
        "NameEN" : "Llangefni", 
        "NameCY" : "Llangefni", 
        "NGR" : "SH4548875758", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Cefni at Llangefni", 
        "TitleCY" : "Afon Cefni at Llangefni", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4292"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.18823383387365, 51.7439134694682
        ]
      }, 
      "properties" : {
        "Location" : "4293", 
        "LatestValue" : "0.393", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Cwm, Elm Street", 
        "NameCY" : "Cwm,Stryd Elm", 
        "NGR" : "SO1805905691", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "River Ebbw at Cwm, Elm Street", 
        "TitleCY" : "Afon Ebwy yn Cwm, Stryd Elm", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4293"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.3143897972761, 53.1656287679909
        ]
      }, 
      "properties" : {
        "Location" : "4294", 
        "LatestValue" : "2.186", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Dwyran Rhyd Y Valley", 
        "NameCY" : "Dwyran Rhyd-Y-Valley", 
        "NGR" : "SH4538165669", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Afon Rhyd Y Valley at Dwyran", 
        "TitleCY" : "Afon Rhyd-Y-Valley yn Nwyran", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4294"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.77999655661545, 51.992977629508
        ]
      }, 
      "properties" : {
        "Location" : "4206", 
        "LatestValue" : "0.957", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Pont Felindre", 
        "NameCY" : "Pont Felindre", 
        "NGR" : "SN7788234221", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Gwydderig at Pont Felindre", 
        "TitleCY" : "Afon Gwydderig ym Mhont Felindre", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4206"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.31017232017473, 51.8539593575695
        ]
      }, 
      "properties" : {
        "Location" : "4207", 
        "LatestValue" : "3.565", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Pothouse Wharf", 
        "NameCY" : "Pothouse Wharf", 
        "NGR" : "SN4099319786", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Tywi at Pothouse Wharf", 
        "TitleCY" : "Afon Tywi yn Pothouse Wharf", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4207"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.19218416629047, 51.8617128827449
        ]
      }, 
      "properties" : {
        "Location" : "4208", 
        "LatestValue" : "2.849", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Ty Castell", 
        "NameCY" : "Ty Castell", 
        "NGR" : "SN4914420397", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Tywi at Ty Castell", 
        "TitleCY" : "Afon Tywi yn Nhy Castell", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4208"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.77477339948646, 52.110127930859
        ]
      }, 
      "properties" : {
        "Location" : "4209", 
        "LatestValue" : "0.460", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Ystradffin", 
        "NameCY" : "Ystradffin", 
        "NGR" : "SN7855947241", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Tywi at Ystradffin", 
        "TitleCY" : "Afon Tywi yn Ystradffin", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4209"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.56175386323039, 52.0454127179634
        ]
      }, 
      "properties" : {
        "Location" : "4210", 
        "LatestValue" : "2.375", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Glan Teifi", 
        "NameCY" : "Glan Teifi", 
        "NGR" : "SN2441941655", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Teifi at Glan Teifi", 
        "TitleCY" : "Afon Teifi yng Nglan Teifi", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4210"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.3173167015523, 52.0365201946406
        ]
      }, 
      "properties" : {
        "Location" : "4211", 
        "LatestValue" : "0.836", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Llandysul", 
        "NameCY" : "Llandysul", 
        "NGR" : "SN4114840103", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Tyweli at Llandysul", 
        "TitleCY" : "Afon Tyweli yn Llandysul", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4211"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.28571329664483, 52.0421016553278
        ]
      }, 
      "properties" : {
        "Location" : "4212", 
        "LatestValue" : "1.440", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Llanfair", 
        "NameCY" : "Llanfair", 
        "NGR" : "SN4333540655", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Teifi at Llanfair", 
        "TitleCY" : "Afon Teifi yn Llanfair", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4212"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.97312303820593, 52.1937356170609
        ]
      }, 
      "properties" : {
        "Location" : "4213", 
        "LatestValue" : "1.605", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Pont Llanio", 
        "NameCY" : "Pont Llanio", 
        "NGR" : "SN6523056890", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Teifi at Pont Llanio", 
        "TitleCY" : "Afon Teifi ym Mhont Llanio", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4213"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.86281014147683, 52.2820278889324
        ]
      }, 
      "properties" : {
        "Location" : "4214", 
        "LatestValue" : "0.857", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Pontrhydfendigaid", 
        "NameCY" : "Pontrhydfendigaid", 
        "NGR" : "SN7302266510", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Teifi at Pontrhydfendigaid", 
        "TitleCY" : "Afon Teifi ym Mhontrhydfendigaid", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4214"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.93334099563108, 52.2203803223271
        ]
      }, 
      "properties" : {
        "Location" : "4216", 
        "LatestValue" : "0.478", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Tregaron", 
        "NameCY" : "Tregaron", 
        "NGR" : "SN6802859780", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Brennig at Tregaron", 
        "TitleCY" : "Afon Brenig yn Nhregaron", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4216"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.44610002329767, 51.4621548777044
        ]
      }, 
      "properties" : {
        "Location" : "4217", 
        "LatestValue" : "0.454", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Cowbridge", 
        "NameCY" : "Y Bont-faen", 
        "NGR" : "SS9963674678", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Thaw at Cowbridge", 
        "TitleCY" : "Afon Ddawan yn y Bont-faen", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4217"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.4150140753905, 51.4354139428283
        ]
      }, 
      "properties" : {
        "Location" : "4218", 
        "LatestValue" : "0.431", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Gigman Bridge", 
        "NameCY" : "Pont Gigman", 
        "NGR" : "ST0173871662", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Thaw at Gigman Bridge", 
        "TitleCY" : "Afon Ddawan ym Mhont Gigman", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4218"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.29210971472949, 51.5749950955312
        ]
      }, 
      "properties" : {
        "Location" : "4219", 
        "LatestValue" : "1.014", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Upper Boat Bridge", 
        "NameCY" : "Pont Glan-bad", 
        "NGR" : "ST1055587027", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Taff at Upper Boat Bridge", 
        "TitleCY" : "Afon Taf ym Mhont Glan-bad", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4219"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.33796098908586, 51.6051254957445
        ]
      }, 
      "properties" : {
        "Location" : "4220", 
        "LatestValue" : "0.771", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Pontypridd, Sion Street", 
        "NameCY" : "Pontypridd, Stryd Sion", 
        "NGR" : "ST0743990435", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Taff at Pontypridd, Sion Street", 
        "TitleCY" : "Afon Taf ym Mhontypridd, Stryd Sion", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4220"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.33210211671385, 51.6521024755485
        ]
      }, 
      "properties" : {
        "Location" : "4221", 
        "LatestValue" : "0.605", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Abercynon", 
        "NameCY" : "Abercynon", 
        "NGR" : "ST0794095652", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Cynon at Abercynon", 
        "TitleCY" : "Afon Cynon yn Abercynon", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4221"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.22212602228144, 51.6778371296206
        ]
      }, 
      "properties" : {
        "Location" : "4222", 
        "LatestValue" : "0.457", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Bargoed", 
        "NameCY" : "Bargoed", 
        "NGR" : "ST1559698381", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Rhymney at Bargoed", 
        "TitleCY" : "Afon Rhymni ym Margod", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4222"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.6617249900991, 52.0811186306836
        ]
      }, 
      "properties" : {
        "Location" : "4224", 
        "LatestValue" : "0.199", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Cardigan Quay", 
        "NameCY" : "Cei Aberteifi", 
        "NGR" : "SN1771045872", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Teifi at Cardigan Quay", 
        "TitleCY" : "Afon Teifi yng Nghei Aberteifi", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4224"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.75756897236294, 51.7804320031736
        ]
      }, 
      "properties" : {
        "Location" : "4225", 
        "LatestValue" : "0.803", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Teddy Bear Bridge", 
        "NameCY" : "Teddy Bear Bridge", 
        "NGR" : "SN7885110546", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Tawe at Teddy Bear Bridge", 
        "TitleCY" : "Afon Tawe yn Teddy Bear Bridge", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4225"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.7679867867421, 51.5913413694961
        ]
      }, 
      "properties" : {
        "Location" : "4226", 
        "LatestValue" : "0.243", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Taibach", 
        "NameCY" : "Tai-bach", 
        "NGR" : "SS7762389535", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ffrwdwyllt at Taibach", 
        "TitleCY" : "Afon Ffrwdwyllt yn Nhai-bach", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4226"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.1455104324261, 51.6990180893329
        ]
      }, 
      "properties" : {
        "Location" : "4227", 
        "LatestValue" : "0.507", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Felinfoel", 
        "NameCY" : "Felin-foel", 
        "NGR" : "SN5182502209", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Lliedi at Felinfoel", 
        "TitleCY" : "Afon Lliedi yn Felin-foel", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4227"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.13489976102129, 51.6829708322472
        ]
      }, 
      "properties" : {
        "Location" : "4228", 
        "LatestValue" : "0.277", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Halfway", 
        "NameCY" : "Halfway", 
        "NGR" : "SN5250600403", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dafen at Halfway", 
        "TitleCY" : "Afon Dafen yn Halfway", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4228"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.05524549228697, 51.6551305064569
        ]
      }, 
      "properties" : {
        "Location" : "4229", 
        "LatestValue" : "3.212", 
        "LatestTime" : "13/03/2019 16:15", 
        "NameEN" : "Pont-y-Cob", 
        "NameCY" : "Pont-y-Cob", 
        "NGR" : "SS5792597149", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Lliw at Pont-y-Cob", 
        "TitleCY" : "Afon Lliw ym Mhont-y-Cob", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4229"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.05763152165099, 52.4033445231202
        ]
      }, 
      "properties" : {
        "Location" : "4230", 
        "LatestValue" : "0.637", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Llanbadarn Fawr", 
        "NameCY" : "Llanbadarn Fawr", 
        "NGR" : "SN6011780362", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Rheidol at Llanbadarn Fawr", 
        "TitleCY" : "Afon Rheidol yn Llanbadarn Fawr", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4230"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.44068969617126, 51.7131335467888
        ]
      }, 
      "properties" : {
        "Location" : "4231", 
        "LatestValue" : "0.355", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Aberdare", 
        "NameCY" : "Aberdâr", 
        "NGR" : "SO0056202582", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Cynon at Aberdare", 
        "TitleCY" : "Afon Cynon yn Aberdâr", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4231"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.14504818961958, 51.7063620581227
        ]
      }, 
      "properties" : {
        "Location" : "4232", 
        "LatestValue" : "0.635", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Aberbeeg", 
        "NameCY" : "Aber-bîg", 
        "NGR" : "SO2097501467", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ebbw at Aberbeeg", 
        "TitleCY" : "Afon Ebwy yn Aber-bîg", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4232"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.39444766792092, 51.9476885322723
        ]
      }, 
      "properties" : {
        "Location" : "4233", 
        "LatestValue" : "0.924", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Brecon", 
        "NameCY" : "Aberhonddu", 
        "NGR" : "SO0425628606", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Usk at Brecon", 
        "TitleCY" : "Afon Wysg yn Aberhonddu", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4233"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.3122045899421, 52.5158896064988
        ]
      }, 
      "properties" : {
        "Location" : "4235", 
        "LatestValue" : "1.150", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Newtown", 
        "NameCY" : "Y Drenewydd", 
        "NGR" : "SO1105391698", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Severn at Newtown", 
        "TitleCY" : "Afon Hafren yn y Drenewydd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4235"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.91122033250936, 53.000865275981
        ]
      }, 
      "properties" : {
        "Location" : "4241", 
        "LatestValue" : "14.280", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Bangor on Dee", 
        "NameCY" : "Bangor is y Coed", 
        "NGR" : "SJ3894645226", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dee at Bangor on Dee", 
        "TitleCY" : "Afon Dyfrdwy am Bont Leadmill, Y Wyddgrug", 
        "Units" : "mAOD", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4241"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.13749725679566, 53.1723952136431
        ]
      }, 
      "properties" : {
        "Location" : "4242", 
        "LatestValue" : "0.489", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Mold Leadmill Bridge", 
        "NameCY" : "Bont Leadmill, Y Wyddgrug", 
        "NGR" : "SJ2406464524", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Alyn at Mold Leadmill Bridge", 
        "TitleCY" : "Afon Alun ym Mhont Leadmill, Y Wyddgrug", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4242"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.08247241668301, 53.1373731035844
        ]
      }, 
      "properties" : {
        "Location" : "4243", 
        "LatestValue" : "0.925", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Pontblyddyn", 
        "NameCY" : "Pontblyddyn", 
        "NGR" : "SJ2768360571", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Alyn at Pontblyddyn", 
        "TitleCY" : "Afon Alun ym Mhontblyddyn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4243"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.95209924947991, 53.1073122857276
        ]
      }, 
      "properties" : {
        "Location" : "4244", 
        "LatestValue" : "0.437", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Rossett Alyn Bridge", 
        "NameCY" : "Bont Alyn, Yr Orsedd", 
        "NGR" : "SJ3636057103", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Alyn at Rossett Alyn Bridge", 
        "TitleCY" : "Afon Alun am Mhont Alyn, Yr Orsedd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4244"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.27495705895706, 52.5360039682458
        ]
      }, 
      "properties" : {
        "Location" : "4259", 
        "LatestValue" : "0.931", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Aberbechan", 
        "NameCY" : "Aberbechan", 
        "NGR" : "SO1362093890", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Bechan Brook at Aberbechan", 
        "TitleCY" : "Nant Bechan yn Aberbechan", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4259"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.84983798165196, 51.9194702238894
        ]
      }, 
      "properties" : {
        "Location" : "4015", 
        "LatestValue" : "0.901", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Grosmont, Rhosllwyn", 
        "NameCY" : "Y Grysmwnt, Rhosllwyn", 
        "NGR" : "SO4165024890", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Monnow at Grosmont, Rhosllwyn", 
        "TitleCY" : "Afon Mynwy yn y Grysmwnt, Rhosllwyn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4015"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.12670208045574, 52.0765573414897
        ]
      }, 
      "properties" : {
        "Location" : "4016", 
        "LatestValue" : "1.474", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Hay On Wye", 
        "NameCY" : "Y Gelli Gandryll", 
        "NGR" : "SO2288042620", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Wye at Hay On Wye", 
        "TitleCY" : "Afon Gwy yn y Gelli Gandryll", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4016"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.72019317627434, 52.0517639606718
        ]
      }, 
      "properties" : {
        "Location" : "4017", 
        "LatestValue" : "2.851", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Old Wye Bridge, Hereford", 
        "NameCY" : "Old Wye Bridge, Hereford", 
        "NGR" : "SO5071239508", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Wye at Old Wye Bridge, Hereford", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Gwy yn Bont Hen y Gwy, Henffordd", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/4017"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.31466813787552, 52.3054087133497
        ]
      }, 
      "properties" : {
        "Location" : "4018", 
        "LatestValue" : "0.633", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Llanddewi", 
        "NameCY" : "Llanddewi", 
        "NGR" : "SO1046068290", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ithon at Llanddewi", 
        "TitleCY" : "Afon Ieithon yn Llanddewi", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4018"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.65999791619024, 52.0619345176754
        ]
      }, 
      "properties" : {
        "Location" : "4019", 
        "LatestValue" : "2.591", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Lugwardine", 
        "NameCY" : "Lugwardine", 
        "NGR" : "SO5485040600", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Lugg at Lugwardine", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Lugg yn Lugwardine", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/4019"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.11295193872118, 52.1514026791906
        ]
      }, 
      "properties" : {
        "Location" : "4020", 
        "LatestValue" : "0.444", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Michaelchurch on Arrow", 
        "NameCY" : "Llanfihangel Dyffryn Arwy", 
        "NGR" : "SO2395050930", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Arrow at Michaelchurch on Arrow", 
        "TitleCY" : "Afon Arwy yn Llanfihangel Dyffryn Arwy", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4020"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.7228314398863, 51.7974109617377
        ]
      }, 
      "properties" : {
        "Location" : "4021", 
        "LatestValue" : "0.741", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Mitchel Troy", 
        "NameCY" : "Llanfihangel Troddi", 
        "NGR" : "SO5025011220", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Trothy at Mitchel Troy", 
        "TitleCY" : "Afon Troddi yn Llanfihangel Troddi", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4021"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.11748510505652, 52.3088780218102
        ]
      }, 
      "properties" : {
        "Location" : "4022", 
        "LatestValue" : "0.435", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Monaughty", 
        "NameCY" : "Monaughty", 
        "NGR" : "SO2391068450", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Lugg at Monaughty", 
        "TitleCY" : "Afon Llugwy yn Monaughty", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4022"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.70324775908361, 51.8146399296267
        ]
      }, 
      "properties" : {
        "Location" : "4023", 
        "LatestValue" : "3.495", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Monmouth", 
        "NameCY" : "Trefynwy", 
        "NGR" : "SO5161913123", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Wye at Monmouth", 
        "TitleCY" : "Afon Gwy yn Nhrefynwy", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4023"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.72011350012918, 51.8088464220434
        ]
      }, 
      "properties" : {
        "Location" : "4024", 
        "LatestValue" : "2.892", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Monnow Gate", 
        "NameCY" : "Porth Trefynwy", 
        "NGR" : "SO5045012490", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Monnow at Monnow Gate", 
        "TitleCY" : "Afon Mynwy ym Mhorth Trefynwy", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4024"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.62988748136107, 52.0264060635522
        ]
      }, 
      "properties" : {
        "Location" : "4025", 
        "LatestValue" : "3.374", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Mordiford", 
        "NameCY" : "Mordiford", 
        "NGR" : "SO5688036630", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Wye at Mordiford", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Gwy yn Mordiford", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/4025"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.7013246320372, 52.4283242645003
        ]
      }, 
      "properties" : {
        "Location" : "4026", 
        "LatestValue" : "0.327", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Pant Mawr", 
        "NameCY" : "Pant Mawr", 
        "NGR" : "SN8442082510", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Wye at Pant Mawr", 
        "TitleCY" : "Afon Gwy ym Mhant Mawr", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4026"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.95552128388676, 52.0413869426644
        ]
      }, 
      "properties" : {
        "Location" : "4027", 
        "LatestValue" : "0.470", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Peterchurch", 
        "NameCY" : "Peterchurch", 
        "NGR" : "SO3456038540", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Dore at Peterchurch", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Dore yn Peterchurch", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/4027"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.68641733491972, 51.7963165609793
        ]
      }, 
      "properties" : {
        "Location" : "4028", 
        "LatestValue" : "3.938", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Redbrook", 
        "NameCY" : "Redbrook", 
        "NGR" : "SO5276011074", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Wye at Redbrook", 
        "TitleCY" : "Afon Gwy yn Redbrook", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4028"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.58759357399661, 51.9153539190761
        ]
      }, 
      "properties" : {
        "Location" : "4029", 
        "LatestValue" : "3.441", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Ross On Wye", 
        "NameCY" : "Ross On Wye", 
        "NGR" : "SO5968224254", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Wye at Ross On Wye", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Gwy yn Rhosan ar Wy", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/4029"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.05403416564727, 51.9587279654672
        ]
      }, 
      "properties" : {
        "Location" : "4031", 
        "LatestValue" : "0.402", 
        "LatestTime" : "13/03/2019 19:30", 
        "NameEN" : "Tafalog", 
        "NameCY" : "Tafalog", 
        "NGR" : "SO2767029440", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Honddu at Tafalog", 
        "TitleCY" : "Afon Honddu yn Nhafalog", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4031"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.21694052282439, 52.0305245735828
        ]
      }, 
      "properties" : {
        "Location" : "4032", 
        "LatestValue" : "0.676", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Three Cocks", 
        "NameCY" : "Three Cocks", 
        "NGR" : "SO1661037600", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Llynfi at Three Cocks", 
        "TitleCY" : "Afon Llynfi yn Three Cocks", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4032"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.74355289673673, 52.0691688355943
        ]
      }, 
      "properties" : {
        "Location" : "4033", 
        "LatestValue" : "0.140", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Three Elms", 
        "NameCY" : "Three Elms", 
        "NGR" : "SO4913041460", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Yazor Brook at Three Elms", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Nant Yazor yn Tri Llwyfen", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/4033"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.98468738530035, 52.2205135673713
        ]
      }, 
      "properties" : {
        "Location" : "4035", 
        "LatestValue" : "0.697", 
        "LatestTime" : "13/03/2019 19:02", 
        "NameEN" : "Titley Mill", 
        "NameCY" : "Titley Mill", 
        "NGR" : "SO3283058490", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Arrow at Titley Mill", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Arrow yn Melin Titley", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/4035"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.94853271076387, 51.7451850327998
        ]
      }, 
      "properties" : {
        "Location" : "4039", 
        "LatestValue" : "1.752", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Chainbridge", 
        "NameCY" : "Pont Gadwyni", 
        "NGR" : "SO3461005590", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Usk at Chainbridge", 
        "TitleCY" : "Afon Wysg ym Mhont Gadwyni", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4039"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.14144448825858, 51.8568060507483
        ]
      }, 
      "properties" : {
        "Location" : "4040", 
        "LatestValue" : "0.756", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Crickhowell", 
        "NameCY" : "Crucywel", 
        "NGR" : "SO2148618195", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Usk at Crickhowell", 
        "TitleCY" : "Afon Wysg yng Nghrucywel", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4040"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.26984919143344, 51.8758552347433
        ]
      }, 
      "properties" : {
        "Location" : "4041", 
        "LatestValue" : "1.485", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Llandetty", 
        "NameCY" : "Llanddeti", 
        "NGR" : "SO1268020460", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Usk at Llandetty", 
        "TitleCY" : "Afon Wysg yn Llanddeti", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4041"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.01826099666075, 51.8121265590792
        ]
      }, 
      "properties" : {
        "Location" : "4042", 
        "LatestValue" : "1.784", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Llanfoist Bridge", 
        "NameCY" : "Pont Llan-ffwyst", 
        "NGR" : "SO2990013100", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Usk at Llanfoist Bridge", 
        "TitleCY" : "Afon Wysg ym Mhont Llan-ffwyst", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4042"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.10408513529358, 51.8515431519571
        ]
      }, 
      "properties" : {
        "Location" : "4043", 
        "LatestValue" : "0.389", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Millbrook", 
        "NameCY" : "Millbrook", 
        "NGR" : "SO2405017570", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Grwyne Fawr at Millbrook", 
        "TitleCY" : "Afon Grwyne Fawr ym Millbrook", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4043"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.89154530630212, 51.7051701376807
        ]
      }, 
      "properties" : {
        "Location" : "4044", 
        "LatestValue" : "0.568", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Olway Inn", 
        "NameCY" : "Olway Inn", 
        "NGR" : "SO3849001090", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Olway Brook at Olway Inn", 
        "TitleCY" : "Nant Olwy yn Olway Inn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4044"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.45177953643398, 51.9631259648504
        ]
      }, 
      "properties" : {
        "Location" : "4045", 
        "LatestValue" : "0.447", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Pont Ar Yscir", 
        "NameCY" : "Pont-ar-ysgir", 
        "NGR" : "SO0035030400", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Yscir at Pont Ar Yscir", 
        "TitleCY" : "Afon Ysgir ym Mhont-ar-ysgir", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4045"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.55993052330926, 51.9171441197754
        ]
      }, 
      "properties" : {
        "Location" : "4046", 
        "LatestValue" : "0.403", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Pont Hen Hafod", 
        "NameCY" : "Pont Hen Hafod", 
        "NGR" : "SN9281025440", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Senni at Pont Hen Hafod", 
        "TitleCY" : "Afon Senni ym Mhont Hen Hafod", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4046"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.96577371270596, 51.6257906847335
        ]
      }, 
      "properties" : {
        "Location" : "4047", 
        "LatestValue" : "0.708", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Ponthir", 
        "NameCY" : "Pont-hir", 
        "NGR" : "ST3324492327", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Lwyd at Ponthir", 
        "TitleCY" : "Afon Lwyd ym Mhont-hir", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4047"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.10290823245613, 51.6070754121534
        ]
      }, 
      "properties" : {
        "Location" : "4048", 
        "LatestValue" : "0.853", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Risca", 
        "NameCY" : "Rhisga", 
        "NGR" : "ST2372090380", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ebbw at Risca", 
        "TitleCY" : "Afon Ebwy yn Rhisga", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4048"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.78868933478086, 51.8817739486529
        ]
      }, 
      "properties" : {
        "Location" : "4049", 
        "LatestValue" : "1.068", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Skenfrith", 
        "NameCY" : "Ynysgynwraidd", 
        "NGR" : "SO4581020650", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Monnow at Skenfrith", 
        "TitleCY" : "Afon Mynwy yn Ynysgynwraidd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4049"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.53312660840343, 51.9540920852396
        ]
      }, 
      "properties" : {
        "Location" : "4050", 
        "LatestValue" : "0.778", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Trallong", 
        "NameCY" : "Trallong", 
        "NGR" : "SN9474029510", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Usk at Trallong", 
        "TitleCY" : "Llan-ffwyst yn Nhrallong", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4050"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.93031817772053, 51.7328326693997
        ]
      }, 
      "properties" : {
        "Location" : "4051", 
        "LatestValue" : "1.056", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Trostrey Weir", 
        "NameCY" : "Cored Trostre", 
        "NGR" : "SO3585004200", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Usk at Trostrey Weir", 
        "TitleCY" : "Afon Wysg yng Nghored Trostre", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4051"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.45517338076806, 52.7599452403766
        ]
      }, 
      "properties" : {
        "Location" : "2003", 
        "LatestValue" : "0.787", 
        "LatestTime" : "13/03/2019 11:15", 
        "NameEN" : "Vyrnwy Weir", 
        "NameCY" : "Cored Efyrnwy", 
        "NGR" : "SJ0190019030", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Vyrnwy at Vyrnwy Weir, Llanwddyn", 
        "TitleCY" : "Afon Efyrnwy yng Nghored Efyrnwy, Llanwddyn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=2003"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.87222012384981, 52.7244922406375
        ]
      }, 
      "properties" : {
        "Location" : "2004", 
        "LatestValue" : "4.589", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Montford", 
        "NameCY" : "Montford", 
        "NGR" : "SJ4119014450", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Severn at Montford", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Hafren yn Montford", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/2004"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.23357610411887, 52.5535243513983
        ]
      }, 
      "properties" : {
        "Location" : "2009", 
        "LatestValue" : "1.726", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Abermule", 
        "NameCY" : "Aber-miwl", 
        "NGR" : "SO1646095790", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Severn at Abermule", 
        "TitleCY" : "Afon Hafren yn Aber-miwl", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=2009"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.54577210461121, 52.4301011430118
        ]
      }, 
      "properties" : {
        "Location" : "2017", 
        "LatestValue" : "0.862", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Rhos Y Pentref", 
        "NameCY" : "Rhos y Pentref", 
        "NGR" : "SN9500082470", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dulas at Rhos Y Pentref, Cwmbelan", 
        "TitleCY" : "Afon Dulas yn Rhos y Pentref, Cwmbelan", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=2017"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.10877928783042, 52.7689939489148
        ]
      }, 
      "properties" : {
        "Location" : "2020", 
        "LatestValue" : "3.388", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Llanymynech", 
        "NameCY" : "Llanymynech", 
        "NGR" : "SJ2529019620", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Vyrnwy at Llanymynech", 
        "TitleCY" : "Afon Efyrnwy yn Llanymynech", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=2020"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.11031354290106, 52.7943312493527
        ]
      }, 
      "properties" : {
        "Location" : "2025", 
        "LatestValue" : "1", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Llanyblodwel", 
        "NameCY" : "Llanyblodwel", 
        "NGR" : "SJ2523022440", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Tanat at Llanyblodwel", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Tanat yn Llanyblodwel", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/2025"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.80330122728246, 52.4060120995355
        ]
      }, 
      "properties" : {
        "Location" : "2030", 
        "LatestValue" : "1.409", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Onibury", 
        "NameCY" : "Onibury", 
        "NGR" : "SO4545078970", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Onny at Onibury", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Onny yn Onibury", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/2030"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.53915397570606, 52.4512260870571
        ]
      }, 
      "properties" : {
        "Location" : "2033", 
        "LatestValue" : "1.473", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Llanidloes", 
        "NameCY" : "Llanidloes", 
        "NGR" : "SN9550084810", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Severn at Llanidloes", 
        "TitleCY" : "Afon Hafren yn Llanidloes", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=2033"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.42729774956555, 52.5143952206642
        ]
      }, 
      "properties" : {
        "Location" : "2034", 
        "LatestValue" : "1.717", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Caersws", 
        "NameCY" : "Caersws", 
        "NGR" : "SO0324091680", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Severn at Caersws", 
        "TitleCY" : "Afon Hafren yng Nghaersws", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=2034"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.25019653316631, 52.707353606471
        ]
      }, 
      "properties" : {
        "Location" : "2035", 
        "LatestValue" : "1.758", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Meifod", 
        "NameCY" : "Meifod", 
        "NGR" : "SJ1563012920", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Vyrnwy at Meifod", 
        "TitleCY" : "Afon Efyrnwy ym Meifod", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=2035"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.04396418050592, 52.3447939509578
        ]
      }, 
      "properties" : {
        "Location" : "2053", 
        "LatestValue" : "0.999", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Knighton", 
        "NameCY" : "Knighton", 
        "NGR" : "SO2898072370", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Teme at Knighton", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Teme yn Tref-y-clawdd", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/2053"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.60069072833352, 52.4681210248754
        ]
      }, 
      "properties" : {
        "Location" : "2054", 
        "LatestValue" : "0.471", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Bryntail", 
        "NameCY" : "Bryn Tail", 
        "NGR" : "SN9136086780", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Clywedog at Bryntail", 
        "TitleCY" : "Afon Clywedog ym Mryntail", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=2054"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.87616506441068, 52.3591040353141
        ]
      }, 
      "properties" : {
        "Location" : "2057", 
        "LatestValue" : "1.040", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Leintwardine", 
        "NameCY" : "Leintwardine", 
        "NGR" : "SO4043073810", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Teme at Leintwardine", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Teme yn Leintwardine", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/2057"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.32175251805764, 52.7045993833102
        ]
      }, 
      "properties" : {
        "Location" : "2060", 
        "LatestValue" : "1.723", 
        "LatestTime" : "13/03/2019 19:30", 
        "NameEN" : "Pont Robert", 
        "NameCY" : "Pontrobert", 
        "NGR" : "SJ1079012700", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Vyrnwy at Pont Robert", 
        "TitleCY" : "Afon Efyrnwy ym Mhontrobert", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=2060"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.04108357116374, 52.7458741089873
        ]
      }, 
      "properties" : {
        "Location" : "2061", 
        "LatestValue" : "5.604", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Llandrinio", 
        "NameCY" : "Llandrinio", 
        "NGR" : "SJ2982016980", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Severn at Llandrinio", 
        "TitleCY" : "Afon Hafren yn Llandrinio", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=2061"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.43376690726623, 52.6766857269736
        ]
      }, 
      "properties" : {
        "Location" : "2062", 
        "LatestValue" : "1.446", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Llanerfyl", 
        "NameCY" : "Llanerfyl", 
        "NGR" : "SJ0316009740", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Banwy at Llanerfyl", 
        "TitleCY" : "Afon Banw yn Llanerfyl", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=2062"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.19668936848293, 52.0456573128299
        ]
      }, 
      "properties" : {
        "Location" : "4014", 
        "LatestValue" : "1.851", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Glasbury", 
        "NameCY" : "Y Clas-ar-Wy", 
        "NGR" : "SO1802739260", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Wye at Glasbury", 
        "TitleCY" : "Afon Gwy yn y Clas-ar-Wy", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4014"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.34988595064013, 52.0906387241628
        ]
      }, 
      "properties" : {
        "Location" : "4013", 
        "LatestValue" : "2.195", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Erwood", 
        "NameCY" : "Erwyd", 
        "NGR" : "SO0761444447", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Wye at Erwood", 
        "TitleCY" : "Afon Gwy yn Erwyd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4013"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.42892429006082, 52.209675784547
        ]
      }, 
      "properties" : {
        "Location" : "4012", 
        "LatestValue" : "1.356", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Disserth", 
        "NameCY" : "Betws Diserth", 
        "NGR" : "SO0246057790", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ithon at Disserth", 
        "TitleCY" : "Afon Ieithon ym Metws Diserth", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4012"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.50330543922833, 52.2963784294356
        ]
      }, 
      "properties" : {
        "Location" : "4011", 
        "LatestValue" : "0.980", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Ddol Farm", 
        "NameCY" : "Fferm Ddôl", 
        "NGR" : "SN9757867536", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Wye at Ddol Farm", 
        "TitleCY" : "Afon Gwy yn Fferm Ddôl", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4011"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.46957644735914, 52.1465080454136
        ]
      }, 
      "properties" : {
        "Location" : "4010", 
        "LatestValue" : "1.046", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Cilmery", 
        "NameCY" : "Cilmeri", 
        "NGR" : "SN9954050820", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Irfon at Cilmery", 
        "TitleCY" : "Afon Irfon yng Nghilmeri", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4010"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.57196514019698, 52.2690683202661
        ]
      }, 
      "properties" : {
        "Location" : "4008", 
        "LatestValue" : "0.810", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Caban Elan", 
        "NameCY" : "Caban Elan", 
        "NGR" : "SN9283064598", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Elan at Caban Elan", 
        "TitleCY" : "Afon Elan yng Nghaban Elan", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4008"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.93273045897402, 52.2768573493728
        ]
      }, 
      "properties" : {
        "Location" : "4007", 
        "LatestValue" : "0.900", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Byton", 
        "NameCY" : "Byton", 
        "NGR" : "SO3646064710", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Lugg at Byton", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Lugg yn Byton", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/4007"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.411779258155, 52.1535979887997
        ]
      }, 
      "properties" : {
        "Location" : "4005", 
        "LatestValue" : "2.217", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Builth Wells", 
        "NameCY" : "Llanfair-ym-Muallt", 
        "NGR" : "SO0351051530", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Wye at Builth Wells", 
        "TitleCY" : "Afon Gwy yn Llanfair-ym-Muallt", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4005"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.96950158763215, 52.0965128006635
        ]
      }, 
      "properties" : {
        "Location" : "4004", 
        "LatestValue" : "2.831", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Bredwardine", 
        "NameCY" : "Bredwardine", 
        "NGR" : "SO3368344684", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Wye at Bredwardine", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Gwy yn Bredwardine", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/4004"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.75222261994926, 52.0450993895226
        ]
      }, 
      "properties" : {
        "Location" : "4002", 
        "LatestValue" : "3.217", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Belmont", 
        "NameCY" : "Belmont", 
        "NGR" : "SO4850838789", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Wye at Belmont", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Gwy yn Belmont", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/4002"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.61804302914531, 52.1011086109713
        ]
      }, 
      "properties" : {
        "Location" : "4001", 
        "LatestValue" : "0.772", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Abernant", 
        "NameCY" : "Abernant", 
        "NGR" : "SN8926945987", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Irfon at Abernant", 
        "TitleCY" : "Afon Irfon yn Abernant", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4001"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.0378880397439, 52.7797042194998
        ]
      }, 
      "properties" : {
        "Location" : "2096", 
        "LatestValue" : "1.243", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Maesbrook", 
        "NameCY" : "Maesbrook", 
        "NGR" : "SJ3009020740", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Flood Channel at Maesbrook", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Channel Llifogydd yn Maesbrook", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/2096"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.16036608067714, 52.3948732684685
        ]
      }, 
      "properties" : {
        "Location" : "2090", 
        "LatestValue" : "0.588", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Dutlas", 
        "NameCY" : "Dutlas", 
        "NGR" : "SO2114078061", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Teme at Dutlas", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Teme yn Dutlas", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/2090"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.15570951349207, 52.6604850935018
        ]
      }, 
      "properties" : {
        "Location" : "2089", 
        "LatestValue" : "0.324", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Welshpool", 
        "NameCY" : "Y Trallwng", 
        "NGR" : "SJ2193007600", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Lledan Brook at Welshpool", 
        "TitleCY" : "Nant Lledan yn y Trallwng", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=2089"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.1668635678099, 52.6018494538365
        ]
      }, 
      "properties" : {
        "Location" : "2072", 
        "LatestValue" : "2.171", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Munlyn", 
        "NameCY" : "Munlyn", 
        "NGR" : "SJ2107001090", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Severn at Munlyn", 
        "TitleCY" : "Afon Hafren ym Munlyn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=2072"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.11579833140952, 52.6724649654377
        ]
      }, 
      "properties" : {
        "Location" : "2068", 
        "LatestValue" : "4.032", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Buttington", 
        "NameCY" : "Tal-y-bont", 
        "NGR" : "SJ2465008890", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Severn at Buttington", 
        "TitleCY" : "Afon Hafren yn Nhal-y-bont", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=2068"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.99433668549114, 52.7357571414149
        ]
      }, 
      "properties" : {
        "Location" : "2067", 
        "LatestValue" : "5.509", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Crew Green", 
        "NameCY" : "Crew Green", 
        "NGR" : "SJ3296015810", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Environment Agency Station monitoring the Severn at Crew Green", 
        "TitleCY" : "Gorsaf Asiantaeth yr Amgylchedd monitro'r Hafren yn Crew Green", 
        "Units" : "m", 
        "url" : "https://flood-warning-information.service.gov.uk/station/2067"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.68339548417901, 51.8261537326102
        ]
      }, 
      "properties" : {
        "Location" : "4091", 
        "LatestValue" : "0.586", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Craig y Nos", 
        "NameCY" : "Craig y Nos", 
        "NGR" : "SN8408515510", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Tawe at Craig y Nos", 
        "TitleCY" : "Afon Tawe yn Nghraig y Nos", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4091"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.7477339276186, 51.6225871403728
        ]
      }, 
      "properties" : {
        "Location" : "4092", 
        "LatestValue" : "0.391", 
        "LatestTime" : "13/03/2019 19:45", 
        "NameEN" : "Cwmafan", 
        "NameCY" : "Cwmafan", 
        "NGR" : "SS7910992976", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Afan at Cwmafan", 
        "TitleCY" : "Afon Afan yng Nghwmafan", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4092"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.9225047978525, 52.3969975086651
        ]
      }, 
      "properties" : {
        "Location" : "4093", 
        "LatestValue" : "1.158", 
        "LatestTime" : "13/03/2019 19:45", 
        "NameEN" : "Cwm Rheidol", 
        "NameCY" : "Cwm Rheidol", 
        "NGR" : "SN6929079403", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Rheidol at Cwm Rheidol", 
        "TitleCY" : "Afon Rheidol yng Nghwm Rheidol", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4093"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.77671281487062, 52.3480026563649
        ]
      }, 
      "properties" : {
        "Location" : "4094", 
        "LatestValue" : "0.437", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Cwmystwyth", 
        "NameCY" : "Cwmystwyth", 
        "NGR" : "SN7907573700", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ystwyth at Cwmystwyth", 
        "TitleCY" : "Afon Ystwyth yng Nghwmystwyth", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4094"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.49617132149611, 51.5207937795814
        ]
      }, 
      "properties" : {
        "Location" : "4095", 
        "LatestValue" : "0.585", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Felindre Road", 
        "NameCY" : "Ffordd Felindre", 
        "NGR" : "SS9629181269", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ewenny Fawr at Felindre Road", 
        "TitleCY" : "Afon Ewenni Fawr yn Ffordd Felindre", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4095"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.48326295583187, 51.8297738461957
        ]
      }, 
      "properties" : {
        "Location" : "4096", 
        "LatestValue" : "0.819", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Glasfryn Ford", 
        "NameCY" : "Rhyd Glasfryn", 
        "NGR" : "SN2898217489", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dewi Fawr at Glasfryn Ford", 
        "TitleCY" : "Afon Dewi Fawr yn Rhyd Glasfryn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4096"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.78029314917681, 51.7688408834458
        ]
      }, 
      "properties" : {
        "Location" : "4097", 
        "LatestValue" : "0.563", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Gurnos", 
        "NameCY" : "Y Gurnos", 
        "NGR" : "SN7725209295", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Twrch at Gurnos", 
        "TitleCY" : "Afon Twrch yn y Gurnos", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4097"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.97017321029636, 51.8044816554656
        ]
      }, 
      "properties" : {
        "Location" : "4098", 
        "LatestValue" : "0.753", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Haverfordwest", 
        "NameCY" : "Hwlffordd", 
        "NGR" : "SM9532215933", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Western Cleddau at Haverfordwest", 
        "TitleCY" : "Afon Cleddau Wen yn Hwlffordd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4098"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.55718765369494, 51.4925552257676
        ]
      }, 
      "properties" : {
        "Location" : "4099", 
        "LatestValue" : "0.411", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Keepers Lodge", 
        "NameCY" : "Keepers Lodge", 
        "NGR" : "SS9199178217", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ewenny at Keepers Lodge", 
        "TitleCY" : "Afon Ewenni yn Keepers Lodge", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4099"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.712328492759, 51.6716244569232
        ]
      }, 
      "properties" : {
        "Location" : "4100", 
        "LatestValue" : "0.575", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Kiln Park", 
        "NameCY" : "Parc Kiln", 
        "NGR" : "SN1254600467", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ritec at Kiln Park", 
        "TitleCY" : "Afon Ritec ym Mharc Kiln", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4100"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.16480908076696, 51.6664431695091
        ]
      }, 
      "properties" : {
        "Location" : "4102", 
        "LatestValue" : "0.635", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Llanelli Tidal", 
        "NameCY" : "Llanw Llanelli ", 
        "NGR" : "SS5038498626", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Llanelli at Llanelli Tidal", 
        "TitleCY" : "Llanw Llanelli yn Llanelli ", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4102"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.24232052701974, 51.8006610637959
        ]
      }, 
      "properties" : {
        "Location" : "4103", 
        "LatestValue" : "0.754", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Llangendeirne", 
        "NameCY" : "Llangyndeyrn", 
        "NGR" : "SN4548313713", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Gwendraeth Fach at Llangendeirne", 
        "TitleCY" : "Afon Gwendraeth Fach yn Llangyndeyrn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4103"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.05656128309726, 52.4340512526554
        ]
      }, 
      "properties" : {
        "Location" : "4104", 
        "LatestValue" : "0.601", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Llangorwen", 
        "NameCY" : "Llangorwen", 
        "NGR" : "SN6028783775", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Clarach at Llangorwen", 
        "TitleCY" : "Afon Clarach yn Llangorwen", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4104"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.1401794439362, 52.3067545855822
        ]
      }, 
      "properties" : {
        "Location" : "4105", 
        "LatestValue" : "0.374", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Llanrhystud", 
        "NameCY" : "Llanrhystud", 
        "NGR" : "SN5418469783", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Wyre at Llanrhystud", 
        "TitleCY" : "Afon Wyre yn Llanrhystud", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4105"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.66593958405323, 51.8790659485793
        ]
      }, 
      "properties" : {
        "Location" : "4106", 
        "LatestValue" : "1.284", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Login", 
        "NameCY" : "Login", 
        "NGR" : "SN1659723415", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Taf at Login", 
        "TitleCY" : "Afon Taf yn Login", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4106"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.64590105450514, 51.605814391654
        ]
      }, 
      "properties" : {
        "Location" : "4107", 
        "LatestValue" : "0.324", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Maesteg", 
        "NameCY" : "Maesteg", 
        "NGR" : "SS8611690947", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Llynfi at Maesteg", 
        "TitleCY" : "Afon Llynfi ym Maesteg", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4107"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.77569161130703, 51.6041829139214
        ]
      }, 
      "properties" : {
        "Location" : "4108", 
        "LatestValue" : "0.964", 
        "LatestTime" : "13/03/2019 19:30", 
        "NameEN" : "Marcroft Weir", 
        "NameCY" : "Cored Marcroft", 
        "NGR" : "SS7712490976", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Afan at Marcroft Weir", 
        "TitleCY" : "Afon Afan yn Nghored Marcroft", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4108"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.97870349628356, 51.7915033950793
        ]
      }, 
      "properties" : {
        "Location" : "4109", 
        "LatestValue" : "0.618", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Merlins Bridge", 
        "NameCY" : "Pont Fadlen", 
        "NGR" : "SM9467514514", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Merlin's Brook at Merlins Bridge", 
        "TitleCY" : "Merlin’s Brook ym Mhont Fadlen", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4109"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -5.18927373879647, 51.886197400083
        ]
      }, 
      "properties" : {
        "Location" : "4110", 
        "LatestValue" : "0.394", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Middle Mill", 
        "NameCY" : "Felinganol", 
        "NGR" : "SM8061825657", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Solfach at Middle Mill", 
        "TitleCY" : "Afon Solfach yn Felinganol", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4110"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.93251625788289, 51.6421888498167
        ]
      }, 
      "properties" : {
        "Location" : "4111", 
        "LatestValue" : "3.584", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Morfa", 
        "NameCY" : "Afon Morfa", 
        "NGR" : "SS6637695478", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Tawe at Morfa", 
        "TitleCY" : "Afon Tawe ym Morfa", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4111"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.80872225681624, 51.6649428370395
        ]
      }, 
      "properties" : {
        "Location" : "4112", 
        "LatestValue" : "1.406", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Neath Tidal", 
        "NameCY" : "Llanw Nedd", 
        "NGR" : "SS7500497789", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Neath at Neath Tidal", 
        "TitleCY" : "Llanw Afon Nedd yn Nedd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4112"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.54100393159705, 51.6033002499126
        ]
      }, 
      "properties" : {
        "Location" : "4113", 
        "LatestValue" : "0.479", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Ogmore Vale", 
        "NameCY" : "Cwm Ogwr", 
        "NGR" : "SS9337490509", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ogwr Fawr at Ogmore Vale", 
        "TitleCY" : "Afon Ogwr Fawr yng Nghwm Ogwr", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4113"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.48404762439329, 51.5220718473104
        ]
      }, 
      "properties" : {
        "Location" : "4114", 
        "LatestValue" : "0.254", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Old Mill", 
        "NameCY" : "Old Mill", 
        "NGR" : "SS9713581394", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ewnenny Fach at Old Mill", 
        "TitleCY" : "Afon Ewnenni Fach yn Old Mill", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4114"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.61886247432994, 51.4768915374551
        ]
      }, 
      "properties" : {
        "Location" : "4115", 
        "LatestValue" : "-2.087", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Penybont", 
        "NameCY" : "Pen-y-bont", 
        "NGR" : "SS8767176568", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ogmore at Penybont", 
        "TitleCY" : "Afon Ogwr ym Mhen-y-bont", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4115"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.9736773584201, 51.7953812454431
        ]
      }, 
      "properties" : {
        "Location" : "4116", 
        "LatestValue" : "0.809", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Pont Amman", 
        "NameCY" : "Pontaman", 
        "NGR" : "SN6398912590", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Amman at Pont Amman", 
        "TitleCY" : "Afon Aman ym Mhontaman", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4116"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.07239242822378, 52.3754553915254
        ]
      }, 
      "properties" : {
        "Location" : "4117", 
        "LatestValue" : "1.542", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Pont Llolwyn", 
        "NameCY" : "Pont Llolwyn", 
        "NGR" : "SN5902477289", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ystwyth at Pont Llolwyn", 
        "TitleCY" : "Afon Ystwyth ym Mhont Llolwyn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4117"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -5.17241486147367, 51.898381957293
        ]
      }, 
      "properties" : {
        "Location" : "4118", 
        "LatestValue" : "0.687", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Pont y Cerbyd", 
        "NameCY" : "Pont y Cerbyd", 
        "NGR" : "SM8183726961", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Solfach at Pont y Cerbyd", 
        "TitleCY" : "Afon Solfach ym Mhont y Cerbyd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4118"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.84884465327097, 51.7184331869287
        ]
      }, 
      "properties" : {
        "Location" : "4119", 
        "LatestValue" : "0.435", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Pontardawe", 
        "NameCY" : "Pontardawe", 
        "NGR" : "SN7238003807", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Upper Clydach at Pontardawe", 
        "TitleCY" : "Afon Clydach Uchaf ym Mhontardawe", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4119"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.02978666298624, 51.7156072935334
        ]
      }, 
      "properties" : {
        "Location" : "4120", 
        "LatestValue" : "0.385", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Pontardulais", 
        "NameCY" : "Pontarddulais", 
        "NGR" : "SN5987303825", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dulais at Pontardulais", 
        "TitleCY" : "Afon Dulais ym Mhontarddulais", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4120"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.57368168187079, 51.7617905434047
        ]
      }, 
      "properties" : {
        "Location" : "4121", 
        "LatestValue" : "0.521", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Pontneddfechan", 
        "NameCY" : "Pontneddfechan", 
        "NGR" : "SN9149108183", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Mellte at Pontneddfechan", 
        "TitleCY" : "Afon Mellte ym Mhontneddfechan", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4121"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.21905404637229, 51.7539634179743
        ]
      }, 
      "properties" : {
        "Location" : "4122", 
        "LatestValue" : "0.492", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Pontyates", 
        "NameCY" : "Pont-iets", 
        "NGR" : "SN4692908471", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Gwendraeth Fawr at Pontyates", 
        "TitleCY" : "Afon Gwendraeth Fawr ym Mhont-iets", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4122"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.18429778930572, 51.7750459577244
        ]
      }, 
      "properties" : {
        "Location" : "4123", 
        "LatestValue" : "0.659", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Pontyberem", 
        "NameCY" : "Pontyberem", 
        "NGR" : "SN4939810743", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Gwendraeth Fawr at Pontyberem", 
        "TitleCY" : "Afon Gwendraeth Fawr ym Mhontyberem", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4123"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -2.90645968746434, 51.7019809139542
        ]
      }, 
      "properties" : {
        "Location" : "4054", 
        "LatestValue" : "1.135", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Usk Town", 
        "NameCY" : "Brynbuga", 
        "NGR" : "SO3745500748", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Usk at Usk Town", 
        "TitleCY" : "Llan-ffwyst ym Mrynbuga", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4054"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.18784111478001, 51.6244996739025
        ]
      }, 
      "properties" : {
        "Location" : "4055", 
        "LatestValue" : "0.566", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Ynysddu", 
        "NameCY" : "Ynys-ddu", 
        "NGR" : "ST1787092410", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Sirhowy at Ynysddu", 
        "TitleCY" : "Afon Sirhywi yn Ynys-ddu", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4055"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.31796921242531, 51.6477661003841
        ]
      }, 
      "properties" : {
        "Location" : "4056", 
        "LatestValue" : "0.769", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Fiddlers Elbow", 
        "NameCY" : "Fiddlers Elbow", 
        "NGR" : "ST0890995152", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Taff at Fiddlers Elbow", 
        "TitleCY" : "Afon Taf yn Fiddlers Elbow", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4056"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.47919995870414, 51.6441277704038
        ]
      }, 
      "properties" : {
        "Location" : "4057", 
        "LatestValue" : "0.258", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Gelli", 
        "NameCY" : "Y Gelli", 
        "NGR" : "SS9774694961", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Rhondda at Gelli", 
        "TitleCY" : "Afon Rhondda yn y Gelli", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4057"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.39470653322665, 51.5347321538581
        ]
      }, 
      "properties" : {
        "Location" : "4058", 
        "LatestValue" : "0.452", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Lanelay", 
        "NameCY" : "Glanelái", 
        "NGR" : "ST0336082680", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ely at Lanelay", 
        "TitleCY" : "Afon Elái yng Nglanelái", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4058"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.48286931839852, 51.676398393082
        ]
      }, 
      "properties" : {
        "Location" : "4059", 
        "LatestValue" : "0.224", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Maerdy", 
        "NameCY" : "Y Maerdy", 
        "NGR" : "SS9756598555", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Rhondda Fach at Maerdy", 
        "TitleCY" : "Afon Rhondda Fach yn y Maerdy", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4059"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.33052802724707, 51.5988347781378
        ]
      }, 
      "properties" : {
        "Location" : "4061", 
        "LatestValue" : "1.059", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Pontypridd", 
        "NameCY" : "Pontypridd", 
        "NGR" : "ST0794189726", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Taff at Pontypridd", 
        "TitleCY" : "Afon Taf ym Mhontypridd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4061"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.26761296132338, 51.4854600997534
        ]
      }, 
      "properties" : {
        "Location" : "4062", 
        "LatestValue" : "0.645", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "St Fagans", 
        "NameCY" : "Sain Ffagan", 
        "NGR" : "ST1208077040", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ely at St Fagans", 
        "TitleCY" : "Afon Elái yn Sain Ffagan", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4062"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.36883031618226, 51.6093103860422
        ]
      }, 
      "properties" : {
        "Location" : "4063", 
        "LatestValue" : "0.664", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Trehafod", 
        "NameCY" : "Trehafod", 
        "NGR" : "ST0531090940", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Rhondda at Trehafod", 
        "TitleCY" : "Afon Rhondda yn Nhrehafod", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4063"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.34592371998524, 51.7087547487603
        ]
      }, 
      "properties" : {
        "Location" : "4064", 
        "LatestValue" : "0.701", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Troedyrhiw", 
        "NameCY" : "Troed-y-rhiw", 
        "NGR" : "SO0710001970", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Taff at Troedyrhiw", 
        "TitleCY" : "Afon Taf yn Nhroed-y-rhiw", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4064"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.54516098612141, 51.6767821317465
        ]
      }, 
      "properties" : {
        "Location" : "4065", 
        "LatestValue" : "0.334", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Tynewydd", 
        "NameCY" : "Tynewydd", 
        "NGR" : "SS9325998687", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Rhondda at Tynewydd", 
        "TitleCY" : "Afon Rhondda yn Nhynewydd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4065"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.20988534441, 51.4973099653968
        ]
      }, 
      "properties" : {
        "Location" : "4067", 
        "LatestValue" : "1.314", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Western Avenue", 
        "NameCY" : "Rhodfa'r Gorllewin", 
        "NGR" : "ST1611078290", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Taff at Western Avenue", 
        "TitleCY" : "Afon Taf ger Rhodfa'r Gorllewin", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4067"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.50683021552129, 51.73809224799
        ]
      }, 
      "properties" : {
        "Location" : "4068", 
        "LatestValue" : "0.446", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Hirwaun", 
        "NameCY" : "Hirwaun", 
        "NGR" : "SN9605005450", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Cynon at Hirwaun", 
        "TitleCY" : "Afon Cynon yn Hirwaun", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4068"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.35349532796462, 51.6045247380961
        ]
      }, 
      "properties" : {
        "Location" : "4069", 
        "LatestValue" : "0.390", 
        "LatestTime" : "03/07/2018 12:00", 
        "NameEN" : "Hopkinstown", 
        "NameCY" : "Trehopcyn", 
        "NGR" : "ST0636290388", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Rhondda at Hopkinstown", 
        "TitleCY" : "Afon Rhondda yn Nhrehopcyn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4069"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.11970814308513, 51.5327344295043
        ]
      }, 
      "properties" : {
        "Location" : "4070", 
        "LatestValue" : "0.700", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Llanedeyrn", 
        "NameCY" : "Llanedern", 
        "NGR" : "ST2243082130", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Rhymney at Llanedeyrn", 
        "TitleCY" : "Afon Rhymni yn Llanedern", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4070"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.16338002850152, 51.5887777951602
        ]
      }, 
      "properties" : {
        "Location" : "4071", 
        "LatestValue" : "0.579", 
        "LatestTime" : "13/03/2019 19:00", 
        "NameEN" : "Waterloo Bridge, Machen", 
        "NameCY" : "Waterloo Bridge, Machen", 
        "NGR" : "ST1950088410", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Rhymney at Waterloo Bridge, Machen", 
        "TitleCY" : "Afon Rhymni yn Waterloo Bridge, Machen", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4071"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.60144719704148, 51.5572514893908
        ]
      }, 
      "properties" : {
        "Location" : "4090", 
        "LatestValue" : "0.448", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Coytrahen", 
        "NameCY" : "Goetre-hen", 
        "NGR" : "SS8907685478", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Llynfi at Coytrahen", 
        "TitleCY" : "Afon Llynfi yng Nghoetre-hen", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4090"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.55719543746814, 51.8151036043277
        ]
      }, 
      "properties" : {
        "Location" : "4089", 
        "LatestValue" : "1.734", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Clog Y Fran", 
        "NameCY" : "Clog y Fran", 
        "NGR" : "SN2383116034", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Taf at Clog Y Fran", 
        "TitleCY" : "Afon Taf yng Nghlog y Fran", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4089"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.90614251952411, 51.9767907884681
        ]
      }, 
      "properties" : {
        "Location" : "4088", 
        "LatestValue" : "0.478", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Cilrhedyn Bridge", 
        "NameCY" : "Pont Cilrhedyn", 
        "NGR" : "SN0050134914", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Gwaun at Cilrhedyn Bridge", 
        "TitleCY" : "Afon Gwaun ym Mhont Cilrhedyn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4088"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.76863410049533, 51.6922783471249
        ]
      }, 
      "properties" : {
        "Location" : "4087", 
        "LatestValue" : "0.497", 
        "LatestTime" : "13/03/2019 19:30", 
        "NameEN" : "Cilfrew", 
        "NameCY" : "Cil-ffriw", 
        "NGR" : "SN7785000761", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Dulais at Cilfrew", 
        "TitleCY" : "Afon Dulais yng Nghil-ffriw", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4087"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.79738367464752, 51.8029579589644
        ]
      }, 
      "properties" : {
        "Location" : "4085", 
        "LatestValue" : "0.987", 
        "LatestTime" : "13/03/2019 19:30", 
        "NameEN" : "Canaston Bridge", 
        "NameCY" : "Pont Canaston", 
        "NGR" : "SN0722615292", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Eastern Cleddau at Canaston Bridge", 
        "TitleCY" : "Afon Cleddau Ddu ym Mhont Canaston", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4085"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.58248966472803, 51.5480400322596
        ]
      }, 
      "properties" : {
        "Location" : "4084", 
        "LatestValue" : "0.606", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Brynmenyn", 
        "NameCY" : "Brynmenyn", 
        "NGR" : "SS9036884425", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ogmore at Brynmenyn", 
        "TitleCY" : "Afon Ogwr ym Mrynmenyn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4084"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.5807772380858, 51.5032100580336
        ]
      }, 
      "properties" : {
        "Location" : "4083", 
        "LatestValue" : "0.686", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Bridgend", 
        "NameCY" : "Pen-y-bont ar Ogwr", 
        "NGR" : "SS9037979437", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ogmore at Bridgend", 
        "TitleCY" : "Afon Ogwr ym Mhen-y-bont ar Ogwr", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4083"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.46142471922831, 51.4078766455988
        ]
      }, 
      "properties" : {
        "Location" : "4082", 
        "LatestValue" : "0.591", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Upper Boverton", 
        "NameCY" : "Trebeferad Uchaf", 
        "NGR" : "SS9845168663", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Boverton Brook at Upper Boverton", 
        "TitleCY" : "Nant Trebeferad yn Nhrebeferad Uchaf", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4082"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.99503075771987, 51.5981766277811
        ]
      }, 
      "properties" : {
        "Location" : "4081", 
        "LatestValue" : "0.713", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Blackpill", 
        "NameCY" : "Dulais", 
        "NGR" : "SS6191790700", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Clyne at Blackpill", 
        "TitleCY" : "Afon Clun yn Nulais", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4081"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.08634618070258, 52.4115178888125
        ]
      }, 
      "properties" : {
        "Location" : "4080", 
        "LatestValue" : "0.439", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Aberystwyth Tidal", 
        "NameCY" : "Llanw Aberystwyth ", 
        "NGR" : "SN5819081327", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Rheidol at Aberystwyth Tidal", 
        "TitleCY" : "Llanw Afon Rheidol yn Aberystwyth ", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4080"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.59182131286845, 51.546807722402
        ]
      }, 
      "properties" : {
        "Location" : "4079", 
        "LatestValue" : "0.716", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Aberkenfig", 
        "NameCY" : "Abercynffig", 
        "NGR" : "SS8971884302", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Llynfi at Aberkenfig", 
        "TitleCY" : "Afon Llynfi yn Abercynffig", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4079"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -4.2572716893136, 52.2378708292182
        ]
      }, 
      "properties" : {
        "Location" : "4078", 
        "LatestValue" : "0.733", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Aberaeron", 
        "NameCY" : "Aberaeron", 
        "NGR" : "SN4596362365", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Aeron at Aberaeron", 
        "TitleCY" : "Afon Aeron yn Aberaeron", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4078"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.07140370539604, 51.5937898497827
        ]
      }, 
      "properties" : {
        "Location" : "4076", 
        "LatestValue" : "0.789", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Rhiwderin", 
        "NameCY" : "Rhiwderyn", 
        "NGR" : "ST2588088870", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Ebbw at Rhiwderin", 
        "TitleCY" : "Afon Ebwy yn Rhiwderyn", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4076"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.05855453517739, 51.7178103725856
        ]
      }, 
      "properties" : {
        "Location" : "4075", 
        "LatestValue" : "0.470", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Pontnewynydd / Abersychan", 
        "NameCY" : "Pontnewynydd / Abersychan", 
        "NGR" : "SO2697002650", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Lwyd at Pontnewynydd / Abersychan", 
        "TitleCY" : "Afon Lwyd ym Mhontnewynydd / Abersychan", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4075"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.24705084761182, 51.7237831401175
        ]
      }, 
      "properties" : {
        "Location" : "4074", 
        "LatestValue" : "0.318", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "New Tredegar", 
        "NameCY" : "Tredegar Newydd", 
        "NGR" : "SO1396003520", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Rhymney at New Tredegar", 
        "TitleCY" : "Afon Rhymni yn Nhredegar Newydd", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4074"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.26061902426698, 51.7951213041503
        ]
      }, 
      "properties" : {
        "Location" : "4073", 
        "LatestValue" : "0.271", 
        "LatestTime" : "13/03/2019 18:00", 
        "NameEN" : "Nant Y Bwch", 
        "NameCY" : "Nant-y-Bwch", 
        "NGR" : "SO1316011470", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Sirhowy at Nant Y Bwch", 
        "TitleCY" : "Afon Sirhywi yn Nant-y-Bwch", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4073"
      }
    }, 
    {
      "type" : "Feature", 
      "geometry" : 
      {
        "type" : "Point", 
        "coordinates" : [
          -3.38762194366533, 51.7517894456581
        ]
      }, 
      "properties" : {
        "Location" : "4072", 
        "LatestValue" : "0.584", 
        "LatestTime" : "13/03/2019 18:45", 
        "NameEN" : "Merthyr Tydfil", 
        "NameCY" : "Merthyr Tudful", 
        "NGR" : "SO0431006810", 
        "ParamNameEN" : "River Level", 
        "ParamNameCY" : "Lefel yr afon", 
        "StatusEN" : "Online", 
        "StatusCY" : "Ar-lein", 
        "TitleEN" : "Taff at Merthyr Tydfil", 
        "TitleCY" : "Afon Taf ym Merthyr Tudful", 
        "Units" : "m", 
        "url" : "http://rloi.naturalresources.wales/ViewDetails?station=4072"
      }
    }
  ]
}`
