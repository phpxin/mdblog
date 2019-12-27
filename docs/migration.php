<?php

$pdo = new PDO("mysql:host=127.0.0.1;dbname=dream128", "root", "lixinxin");
$pdo->query("set names utf8");
$statement = $pdo->query("select * from article");
if (!$statement || $statement->rowCount()<=0){
    echo "no rows";
    exit() ;
}

$data = [] ;
$rows = $statement->fetchAll(PDO::FETCH_ASSOC) ;
foreach($rows as $row){
    $data[trim($row['title'])] = $row;
}

echo 'old art count ', count($data) , PHP_EOL ;

$pdo = new PDO("mysql:host=127.0.0.1;dbname=mdblog", "root", "lixinxin") ;
$pdo->query("set names utf8");
// $statement = $pdo->query("select * from docs");
// if (!$statement || $statement->rowCount()<=0){
//     echo "doc no rows";
//     exit() ;
// }
// echo $statement->rowCount() , PHP_EOL ;
// $rows = $statement->fetchAll(PDO::FETCH_ASSOC) ;
//
// echo 'new art count ', count($rows) , PHP_EOL ;

$now = time();
foreach($rows as $row){
    echo $row['title'] ,PHP_EOL;
    $t = trim($row['title']);

    $statement = $pdo->query("select * from docs where title='$t'");

    if ($statement && $statement->rowCount()>0) {
        $d = $statement->fetch(PDO::FETCH_ASSOC);
        $sql = "select * from click where hash='".$d["hash"]."'" ;
        $statement = $pdo->query($sql) ;
        if ($statement && $statement->rowCount()>0) {
            echo "update",PHP_EOL;
            $sql = "update click set counter=counter+".$row['hits']." where hash='".$d["hash"]."'" ;
            $pdo->exec($sql) ;
        }else{
            echo "insert",PHP_EOL;
            $sql = "insert into click(hash,counter,created_at,updated_at) values('".$d["hash"]."', ".$row['hits'].", ".$now.", 0)" ;
            echo $sql , PHP_EOL ;
            $pdo->exec($sql) ;
        }
    }else{
        echo "{$t} not found" ,PHP_EOL ;
    }
}