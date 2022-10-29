<?php

$output = fopen("temp/output.csv", "w");
$input = fopen("temp/file_names", "r");
$data = [];
$data[] = ["name", "status_code", "status_date"];
while(!feof($input)) {
    $fileName = fgets($input);
    if ($fileName === false) {
        continue;
    }
    $fileName = trim($fileName);
    $xmlObj = simplexml_load_file($fileName);
    $tradeMark = $xmlObj
        ->TradeMarkTransactionBody
        ->TransactionContentDetails
        ->TransactionData
        ->TradeMarkDetails
        ->TradeMark;
    $markFeature = (string) $tradeMark->MarkFeature;
    $name = (string) $tradeMark->WordMarkSpecification->MarkVerbalElementText;
    $currentStatusCode = (string) $tradeMark->MarkCurrentStatusCode;
    $currentStatusDate = (string) $tradeMark->MarkCurrentStatusDate;
    if ($markFeature !== "Word") {
        continue;
    }
    $data[] = [$name, $currentStatusCode, $currentStatusDate];
}
foreach ($data as $row) {
    fputcsv($output, $row);
}
fclose($input);
fclose($output);