<?php

include 'backend.php';

// Let's pretend Wayne Laubscher (id = 7) is logged in
$specific_member_id = 7;
$specific_members_sql = '
  SELECT id, member_type_id, lft, rght, first_name, last_name
  FROM tree_example.members
  WHERE lft >= (SELECT lft FROM tree_example.members WHERE id = ' . $specific_member_id . ')
  AND rght <= (SELECT rght FROM tree_example.members WHERE id = ' . $specific_member_id . ')
  ORDER BY lft
';
$specific_members_with_lefts_and_rights = getNodesFromDatabase($specific_members_sql);
$specific_members_with_children_and_parents = unserializeFromDatabase($specific_members_with_lefts_and_rights);
$specific_members_array = $specific_members_with_children_and_parents->to_array(4);

?>


<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <script src="static/js/external/jquery/jquery-2.0.3.min.js"></script>
    <script>
        $(document).ready(function () {
            $('.accordion .head').click(function () {
                $(this).next().toggle('slow');
                return false;
            }).next().hide();
        });
    </script>
    <style>
        .head, .content {
            border: 1px solid black;
        }

        .head {
            margin-bottom: 0px;
            cursor: pointer;
        }

        .content {
            display: none;
        }
    </style>
</head>
<body>

<div class="accordion">
    <?php
    foreach ($specific_members_array as $level => $members_at_level) {
        echo '<h3 class="head">Level ' . $level . '</h3>';
        echo '<div class="content">';

        foreach ($members_at_level as $member) {
            echo '<div>' . $member->name . '</div>';
        }

        echo '</div>';
    }
    ?>
</div>


</body>
</html>