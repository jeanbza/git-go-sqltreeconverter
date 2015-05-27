<?php

include 'backend.php';

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
limitChildrenToDepth($specific_members_with_children_and_parents, 4);
$specific_members_json = $specific_members_with_children_and_parents->to_json();

?>

<html>
<head>
    <link rel="stylesheet" type="text/css" href="static/css/app.css">

    <script src="static/js/external/jquery/jquery-2.0.3.min.js"></script>
    <script src="static/js/external/jquery-ui/jquery-ui.js"></script>
    <script src="static/js/external/bootstrap/bootstrap-3.3.4.js"></script>
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