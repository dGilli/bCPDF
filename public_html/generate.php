<?php

$name = escapeshellarg($_POST['name']);
$email = escapeshellarg($_POST['email']);
$phone = escapeshellarg($_POST['phone']);

// Path to Go binary
$command = " /usr/local/bin/generate_pdf $name $email $phone";
$output = shell_exec($command);

if (file_exists("./output/business_card.pdf")) {
    header('Content-Type: application/pdf');
    header('Content-Disposition: attachment; filename="business_card.pdf"');
    readfile("../output/business_card.pdf");
    unlink("../output/business_card.pdf"); // Delete after download
} else {
    echo "Error generating PDF.";
}
?>
