<?php

class Member {
    public $id;
    public $member_type;
    public $lft;
    public $rgt;
    public $name;

    function __construct($id, $member_type, $lft, $rgt, $first_name, $last_name) {
        $this->id = $id;
        $this->member_type = $member_type;
        $this->lft = $lft;
        $this->rgt = $rgt;
        $this->name = $first_name . " " . $last_name;
    }
}

function getNodesFromDatabase() {
    $members = array();

    $servername = "localhost";
    $username = "root";
    $password = "";

    // Create connection
    $conn = new mysqli($servername, $username, $password);

    $sql = "SELECT id, member_type_id, lft, rght, first_name, last_name FROM tree_example.members ORDER BY lft";
    $result = $conn->query($sql);

    while ($row = $result->fetch_assoc()) {
        array_push($members, new Member($row["id"], $row["member_type_id"], $row["lft"], $row["rght"], $row["first_name"], $row["last_name"]));
    }

    $conn->close();

    return $members;
}

$members = getNodesFromDatabase();
print_r($members);