<?php

$pdo = new PDO("mysql:host=127.0.0.1;dbname=dream128", "root", "lixinxin");
$pdo->query("set names utf8");
$statement = $pdo->query("select * from article");
if ($statement->rowCount()<=0){
    echo "no rows";
    exit() ;
}

$data = [] ;
while($row = $statement->fetch(PDO::FETCH_ASSOC)){
    $data[$row['title'] = $row;
}

$pdo2 = new PDO("mysql:host=127.0.0.1;dbname=mdblog", "root", "lixinxin") ;
$pdo->query("set names utf8");
$statement = $pdo->query("select * from docs");
if ($statement->rowCount()<=0){
    echo "doc no rows";
    exit() ;
}

while($row = $statement->fetch(PDO::FETCH_ASSOC)){
    if (isset($data[$row['title']) {
        $sql = "select * from click "
    }
}