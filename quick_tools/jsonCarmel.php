<?php

$a = '{
    "code":0,
    "message":"success",
    "data":{
        "total":1,
        "items":[
            {
                "id":1,
                "uniquekey":"1",
                "name":"Lv.1 城市",
                "growth_start":0,
                "bg_color":"#df74e3",
                "img_url":"https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/44d584d0c69882e19bf91ad2364739c8.png",
                "home_img_url":"https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/f40a2ae1b88467df859f8e3be74dff78.png"
            }
        ]
    }
}';

run($a);

/**
 * @param $a
 * @return false|string
 */
function run($a)
{
    $a = json_decode($a, true);

    dfs($a);

    $s = json_encode($a, JSON_UNESCAPED_UNICODE);
    file_put_contents('/tmp/tmp.json', $s);
}


function dfs(&$val)
{
    foreach ($val as $k => &$v) {
        if (is_array($v)) {
            dfs($v);
        }
        if (strpos($k, '_')) {
            $nk       = camelize($k);
            $val[$nk] = $v;
            unset($val[$k]);
        }
    }
}


function camelize($words, $separator = '_')
{
    $words = $separator . str_replace($separator, " ", strtolower($words));
    return ltrim(str_replace(" ", "", ucwords($words)), $separator);
}


