<?php

// Note: this should be an api that your javascript consumes with AJAX

$databaseAndTable = 'converter.members';

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

        $json .= '"id":' . $this->id . ',"member_type":' . $this->member_type . ',"name":"' . $this->name . '","lft":' . $this->lft . ',"rgt":' . $this->rgt . ',"children":[';

        foreach ($this->children as $index => $child) {
            if ($index != 0) {
                $json .= ',';
            }

            $json .= $child->to_json();
        }

        return $json . ']}';
    }

    function to_array($levels = 0) {
        $members_at_each_level = array();
        $members_at_this_level = array($this); // start by processing root
        $members_at_next_level = array();

        // Levels+1 because otherwise we'd never get to the root element
        for ($i = 0; $i < $levels + 1; $i++) {
            array_push($members_at_each_level, array());

            foreach ($members_at_this_level as $member) {
                array_push($members_at_each_level[$i], $member);
                $members_at_next_level = array_merge($members_at_next_level, $member->children);
            }

            $members_at_this_level = $members_at_next_level;
            $members_at_next_level = array();
        }

        return $members_at_each_level;
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

            if ($node->lft != $node->rgt - 1) {
                $node = $node->children[sizeof($node->children) - 1];
            }
        }
    }

    return $root;
}

function limitChildrenToDepth(&$member, $limit = 4, $depth = 0) {
    if ($depth == $limit) {
        // throw away all children below depth 3
        $member->children = array();
    } else {
        for ($i = 0; $i < sizeof($member->children); $i++) {
            limitChildrenToDepth($member->children[$i], $depth + 1);
        }
    }
}

?>