<?php

class MemberWithLeftsAndRights {
    public $id;
    public $member_type;
    public $name;
    public $lft;
    public $rgt;

    function __construct($id, $member_type, $lft, $rgt, $first_name, $last_name) {
        $this->id = $id;
        $this->member_type = $member_type;
        $this->lft = $lft;
        $this->rgt = $rgt;
        $this->name = $first_name . " " . $last_name;
    }
}

class MemberWithChildrenAndParents {
    public $id;
    public $member_type;
    public $name;

    public $lft;
    public $rgt;

    public $parent;
    public $children;

    function __construct($id, $member_type, $name, $lft, $rgt, &$parent, &$children) {
        $this->id = $id;
        $this->member_type = $member_type;
        $this->name = $name;

        $this->lft = $lft;
        $this->rgt = $rgt;

        $this->parent = $parent;
        $this->children = $children;
    }

    function to_json() {
        $json = '{';

        $json .= '"id":'.$this->id.',"member_type":'.$this->member_type.',"name":"'.$this->name.'","lft":'.$this->lft.',"rgt":'.$this->rgt.',"children":[';

        foreach ($this->children as $index => $child) {
            if ($index != 0) {
                $json .= ',';
            }

            $json .= $child->to_json();
        }

        return $json . ']}';
    }
}

// Convert left/right objects to parent/children objects
function unserializeFromDatabase($membersWithLeftsAndRights) {
    $root = null;
    $node = null;

    $foo = 0; // REMOVE THIS

    while (0 < sizeof($membersWithLeftsAndRights) && $foo < 100) {
        $foo++; // REMOVE THIS

        $member = $membersWithLeftsAndRights[0];
        $membersWithLeftsAndRights = array_slice($membersWithLeftsAndRights, 1, sizeof($membersWithLeftsAndRights));

        if ($root == null) {
            $root = new MemberWithChildrenAndParents($member->id, $member->member_type, $member->name, $member->lft, $member->rgt, $tmp = null, $array = array());
            $node = $root;
        } else {
            // REMOVE THIS
            if ($node->rgt == null) {
                return;
            }

            // REMOVE NULLCHECK
            while ($node != null && $member->lft > $node->rgt && $foo < 100) {
                $foo++; // REMOVE THIS
                $node = $node->parent;
            }

            array_push($node->children, new MemberWithChildrenAndParents($member->id, $member->member_type, $member->name, $member->lft, $member->rgt, $node, $array = array()));

            if ($node->lft != $node->rgt-1) {
                $node = $node->children[sizeof($node->children)-1];
            }
        }
    }

    return $root;
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
        array_push($members, new MemberWithLeftsAndRights($row["id"], $row["member_type_id"], $row["lft"], $row["rght"], $row["first_name"], $row["last_name"]));
    }

    $conn->close();

    return $members;
}

$members_with_lefts_and_rights = getNodesFromDatabase();
$members_with_children_and_parents = unserializeFromDatabase($members_with_lefts_and_rights);

$members_json = $members_with_children_and_parents->to_json();

?>

<html>
<head>
    <link rel="stylesheet" type="text/css" href="static/css/app.css">

    <script src="static/js/external/jquery/jquery-2.0.3.min.js"></script>
    <script src="static/js/external/jquery-ui/jquery-ui.js"></script>
    <script src="static/js/external/bootstrap/bootstrap.min.js"></script>
    <script src="static/js/external/d3/d3.v3.min.js"></script>
    <script src="static/js/conditions_tree.js"></script>

    <script type="text/javascript">
        var frontData = JSON.parse('<?php echo $members_json; ?>');
        var treeData = frontData;
        var matchingUsers = frontData['matchingUsers'];
        $(document).ready(function () {
            initTree(treeData);
        });
    </script>
</head>
<body>
    <div id="d3-tree">
    </div>
</body>
</html>