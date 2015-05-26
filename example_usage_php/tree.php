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

function getNodesFromDatabase($sql) {
    $members = array();

    $servername = "localhost";
    $username = "root";
    $password = "";

    // Create connection
    $conn = new mysqli($servername, $username, $password);

    $result = $conn->query($sql);

    while ($row = $result->fetch_assoc()) {
        array_push($members, new MemberWithLeftsAndRights($row["id"], $row["member_type_id"], $row["lft"], $row["rght"], $row["first_name"], $row["last_name"]));
    }

    $conn->close();

    return $members;
}

// Convert left/right objects to tree with parent+child relationships
// Note: data ordered by left (in sql: ORDER BY lft) is very important for this to work
// Note: this runs in nlogn, which is far better than the n^2 if we were using parent_id in database
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

function limitChildrenToFourDepth(&$member, $depth = 0) {
    if ($depth == 3) {
        // throw away all children below depth 3
        $member->children = array();
    } else {
        for ($i = 0; $i < sizeof($member->children); $i++) {
            limitChildrenToFourDepth($member->children[$i], $depth+1);
        }
    }
}

// Note: ORDER BY lft is very important
$all_members_sql = 'SELECT id, member_type_id, lft, rght, first_name, last_name FROM tree_example.members ORDER BY lft';
$all_members_with_lefts_and_rights = getNodesFromDatabase($all_members_sql);
$all_members_with_children_and_parents = unserializeFromDatabase($all_members_with_lefts_and_rights);
$all_members_json = $all_members_with_children_and_parents->to_json();

// Let's pretend Wayne Laubscher (id = 7) is logged in
$specific_member_id = 7;
$specific_members_sql = '
  SELECT id, member_type_id, lft, rght, first_name, last_name
  FROM tree_example.members
  WHERE lft >= (SELECT lft from tree_example.members WHERE id = ' . $specific_member_id . ')
  AND rght <= (SELECT rght from tree_example.members WHERE id = ' . $specific_member_id . ')
  ORDER BY lft
';
$specific_members_with_lefts_and_rights = getNodesFromDatabase($specific_members_sql);
$specific_members_with_children_and_parents = unserializeFromDatabase($specific_members_with_lefts_and_rights);
limitChildrenToFourDepth($specific_members_with_children_and_parents);
$specific_members_json = $specific_members_with_children_and_parents->to_json();

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
        var treeDataAllUsers = JSON.parse('<?php echo $all_members_json; ?>');;
        var treeDataSpecificUsers = JSON.parse('<?php echo $specific_members_json; ?>');;
        $(document).ready(function () {
            initTree(treeDataAllUsers, "d3-tree-all");
            initTree(treeDataSpecificUsers, "d3-tree-single");
        });
    </script>
</head>
<body>
    <h1>All Members</h1>
    <div id="d3-tree-all"></div>
    <h1>Viewing As Wayne Laubscher (id=7)</h1>
    <div id="d3-tree-single"></div>
</body>
</html>