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

    function __construct($id, $member_type, $name, $rgt, $lft, &$parent, &$children) {
        $this->id = $id;
        $this->member_type = $member_type;
        $this->name = $name;

        $this->rgt = $rgt;
        $this->lft = $lft;

        $this->parent = $parent;
        $this->children = $children;
    }
}

// Convert left/right objects to parent/children objects
function unserializeFromDatabase($membersWithLeftsAndRights) {
    $root = null;
    $node = null;
    $null = null;

    $foo = 0; // REMOVE THIS

    while (0 < sizeof($membersWithLeftsAndRights) && $foo < 100) {
        $foo++; // REMOVE THIS

        $member = $membersWithLeftsAndRights[0];
        $membersWithLeftsAndRights = array_slice($membersWithLeftsAndRights, 1, sizeof($membersWithLeftsAndRights));

        if ($root == null) {
            $array = array();
            $root = new MemberWithChildrenAndParents($member->id, $member->name, $member->name, $member->lft, $member->rgt, $null, $array);
            $node = $root;
        } else {
            print_r($node);

            // REMOVE THIS
            if ($node->rgt == null) {
                return;
            }

            // REMOVE NULLCHECK
            while ($node != null && $member->lft > $node->rgt && $foo < 100) {
                $foo++; // REMOVE THIS
                $node = $node->parent;
            }

            $array = array();
            array_push($node->children, new MemberWithChildrenAndParents($member->id, $member->name, $member->name, $member->lft, $member->rgt, $node, $array));

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
$members_with_parents_and_children = unserializeFromDatabase($members_with_lefts_and_rights);
print_r($members_with_parents_and_children);

?>

<html>
<head>
    <script src="static/js/external/jquery/jquery-2.0.3.min.js"></script>
    <script src="static/js/external/jquery-ui/jquery-ui.js"></script>
    <script src="static/js/external/bootstrap/bootstrap.min.js"></script>
    <script src="static/js/external/d3/d3.v3.min.js"></script>
    <script src="static/js/conditions_tree.js"></script>
    <script src="static/js/app.js"></script>

    <script type="text/javascript">
        var frontData = JSON.parse('<?php echo json_encode($members); ?>');
        var treeData = frontData;
        var matchingUsers = frontData['matchingUsers'];
        $(document).ready(function () {
            initTree(treeData);
            rematchUsers(matchingUsers);
        });
    </script>
</head>
<body>
    <div id="d3-tree">
    </div>
</body>
</html>