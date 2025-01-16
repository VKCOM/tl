<?php


use VK\TL;

require_once(__DIR__."/../../gen/cases_php/VK/TL/tl_streams.php");

class test_case {
    /** @var string */
    public $serializing_data = "";

    /** @var string */
    public $data_type = "";

    /** @var string|false */
    public $tl_data = false;

    /**
     * @param mixed $data
     * @param string $data_type
     * @param array|false $tl_data - array of bytes
     */
    public function __construct(mixed $data, string $data_type, $tl_data = false)
    {
        $this->serializing_data = $data;
        $this->tl_data = $tl_data;
        $this->data_type = $data_type;
    }

    public function test_read_write_interaction(): bool
    {
        $w = new TL\tl_output_stream("");
        switch ($this->data_type) {
            case "string": {
                $w->write_string($this->serializing_data);
                break;
            }
            case "int32": {
                $w->write_int32($this->serializing_data);
                break;
            }
            case "uint32": {
                $w->write_uint32($this->serializing_data);
                break;
            }
            case "float": {
                $w->write_float($this->serializing_data);
                break;
            }
            case "double": {
                $w->write_double($this->serializing_data);
                break;
            }
            case "bool": {
                $w->write_bool($this->serializing_data, 0, 1);
                break;
            }
        }
        $serialized_data = $w->get_data();
        if ($this->tl_data && $serialized_data != test_case::binary_to_string($this->tl_data)) {
            return false;
        }
        $r = new TL\tl_input_stream($w->get_data());
        $return_value = "";
        $success = true;
        switch ($this->data_type) {
            case "string": {
                [$return_value, $success] = $r->read_string();
                break;
            }
            case "int32": {
                [$return_value, $success] = $r->read_int32();
                break;
            }
            case "uint32": {
                [$return_value, $success] = $r->read_uint32();
                break;
            }
            case "float": {
                [$return_value, $success] = $r->read_float();
                break;
            }
            case "double": {
                [$return_value, $success] = $r->read_double();
                break;
            }
            case "bool": {
                [$return_value, $success] = $r->read_bool(0, 1);
                break;
            }
        }
        return $success && ($return_value == $this->serializing_data);
    }

    static function binary_to_string(array $binary): string
    {
        $s = "";
        foreach ($binary as $byte) {
            $s .= chr($byte);
        }
        return $s;
    }
}
function bool_to_string(bool $x): string
{
    if ($x) {
        return "true";
    } else {
        return "false";
    }
}

$tests = array(
    "test-int32" => new test_case($data=123, $data_type="int32", $tl_data=[123, 0, 0, 0]),
    "test-uint32" => new test_case($data=((1 << 8) + (1 << 16)), $data_type="int32", $tl_data=[0, 1, 1, 0]),
    "test-bool-true" => new test_case($data=true, $data_type="bool", $tl_data=[1, 0, 0, 0]),
    "test-bool-false" => new test_case($data=false, $data_type="bool", $tl_data=[0, 0, 0, 0]),
    "test-small-string" => new test_case($data=str_repeat("B", 12), $data_type="string", $tl_data=false),
    "test-big-string" => new test_case($data=str_repeat("A", 255), $data_type="string", $tl_data=false),
    "test-huge-string" => new test_case($data=str_repeat("A", (1 << 24)), $data_type="string", $tl_data=false),
);

$failed_tests = array();

foreach ($tests as $test_name => $test) {
    $test_result = $test->test_read_write_interaction();
//    println("Test \"" . $test_name . "\": " . bool_to_string($test_result));
    if (!$test_result) {
        $failed_tests[] = $test_name;
    }
}

if (count($failed_tests) == 0) {
    printf("\033[01;32mAll tests are passed!\033[0m\n");
} else {
    printf("Failed %d tests:\n", count($failed_tests));
    foreach ($failed_tests as $failed_test) {
        printf("\t%s\n", $failed_test);
    }
}
?>