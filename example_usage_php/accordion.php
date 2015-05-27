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
$members_array_counts = array();

foreach ($specific_members_array as $level => $level_members) {
    if ($level + 1 != sizeof($specific_members_array)) {
        $next_level_members = $specific_members_array[$level + 1];

        // Set up IBO members and names
        $next_level_ibo_members = array_filter($next_level_members, function($member) {
            return $member->member_type == 1;
        });
        $next_level_ibo_names = array();
        foreach ($next_level_ibo_members as $member) { array_push($next_level_ibo_names, $member->name); }

        // Set up student members and names
        $next_level_student_members = array_filter($next_level_members, function($member) {
            return $member->member_type == 2;
        });
        $next_level_student_names = array();
        if ($level == 0) {
            foreach ($next_level_student_members as $member) { array_push($next_level_student_names, $member->name); }
        }

        array_push($members_array_counts, array(
            'ibo_count' => sizeof($next_level_ibo_members),
            'ibo_names' => $next_level_ibo_names,
            'student_count' => sizeof($next_level_student_members),
            'student_names' => $next_level_student_names
        ));
    }
}

?>


<!doctype html>
<html lang="en">
<head>
    <script src="http://code.jquery.com/jquery-2.1.4.js"></script>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap-theme.min.css">
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min.js"></script>
    <script>
        $(document).ready(function () {
            $('.accordion .head').click(function () {
                $(this).next().toggle('fast');
                return false;
            });

            $(".popover-names").popover({
                content: function() {
                    names = JSON.parse($(this).attr('data-names'));

                    return names.join('<br>');
                },
                html: true,
                trigger: 'hover',
                placement: 'right'
            });
        });
    </script>
    <style>
        .accordion {
            width: 50%;
            margin-left: 25%;
            margin-top: 5%;
        }

        .head, .content {
            border: 1px solid black;
        }

        .head {
            margin-bottom: 0px;
            cursor: pointer;
        }

        .content {
            display: block;
        }
    </style>
</head>
<body>

<h3>Viewing As Wayne Laubscher (id=7)</h3>

<div class="accordion">
    <?php
    foreach ($members_array_counts as $level => $level_info) {
        $ibo_names = htmlentities(json_encode($level_info['ibo_names']));
        $student_names = htmlentities(json_encode($level_info['student_names']));

        echo '<h3 class="head">Level ' . $level . '</h3>';
        echo '<div class="content">';

        echo '<div><button class="popover-names" data-names="' . $ibo_names . '">IBO: ' . $level_info['ibo_count'] . '</button></div>';
        echo '<div><button class="popover-names" data-names="' . $student_names . '">Students: ' . $level_info['student_count'] . '</button></div>';

        echo '</div>';
    }
    ?>
</div>

</body>
</html>