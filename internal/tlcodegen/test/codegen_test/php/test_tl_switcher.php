<?php

require_once(__DIR__."/../../gen/cases_php/VK/TL/tl_switcher.php");

function check_chance($ns, $function, $percent, $return_empty = false) {
    echo "[INFO] namespace=\"$ns\" function=\"$function\" percent=$percent%\n";

    $NAMESPACE = $ns;
    $FUNCTION = $function;

    $QUERY = $FUNCTION;
    if ($NAMESPACE !== "") {
        $QUERY = $NAMESPACE . "." . $FUNCTION;
    }

    $NAMESPACE_KEY = $NAMESPACE;
    if ($NAMESPACE_KEY === "") {
        $NAMESPACE_KEY = "_common";
    }

    \VK\TL\tl_switcher::$tl_namespaces_info[$NAMESPACE_KEY] = 1;
    \VK\TL\tl_switcher::$tl_namespaces_info[$NAMESPACE_KEY."_percent"] = $percent;

    $n = 0;
    $N = 10000;
    $eps = 0.02;

    $CALLS = 10;

    for ($i = 0; $i < $N; $i++) {
        $mode = \VK\TL\tl_switcher::tl_get_namespace_methods_mode($NAMESPACE_KEY);
        if ($mode == 1) {
            $n += 1;
        }

        // check indepotentancy
        for ($c = 0; $c < $CALLS; $c++) {
            $selected_mode = \VK\TL\tl_switcher::tl_get_query_mode($QUERY);
            if ($selected_mode != $mode) {
                if ($selected_mode == -1 && $return_empty) {
                    continue;
                }
                echo "[ERROR]";
                echo " selected mode for query differ from expected mode for namespace (fail idempotency).\n";
                exit(1);
            }
        }

        \VK\TL\tl_switcher::tl_reset_selected_modes();
    }

    if (abs($n / $N - ($percent / 100)) > $eps) {
        echo "[ERROR] diff between expected percent " . ($percent / 100);
        echo " and actual " . $n / $N . " is " . abs($n / $N - ($percent / 100));
        echo " (more then eps = " . $eps . ")\n";
        exit(1);
    } else {
        echo "[OK] expected chance of selection: $percent%.\n";
    }
}


echo "\n>>>> TEST TL_SWITCHER\n\n";

check_chance("", "", 100, true);

check_chance("test", "query", 20);
check_chance("test", "query", 100);
check_chance("", "query", 20);
check_chance("", "query", 100);

