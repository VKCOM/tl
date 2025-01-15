<?php

require_once(__DIR__."/../../gen/cases_php/VK/TL/tl_streams.php");
require_once(__DIR__."/../../gen/cases_php/VK/TL/factory.php");

function bool_to_str($value) {
    if ($value) {
        return "true";
    }
    return "false";
}

function digit_to_value(string $char): int {
    if (ord('0') <= ord($char) && ord($char) <= ord('9')) {
        return ord($char) - ord('0');
    } else {
        return ord($char) - ord('a') + 10;
    }
}

function value_to_digit(int $value): string {
    if (0 <= $value && $value <= 9) {
        return chr($value + ord('0'));
    } else {
        return chr($value + ord('a') - 10);
    }
}

function hex_to_string(string $data): string {
    $result = "";
    $chunks = explode(" ", $data);
    foreach($chunks as $chunk) {
        for($i = 7; $i >= 0; $i -= 2) {
            $digit = digit_to_value($chunk[$i]) + 16 * digit_to_value($chunk[$i - 1]);
            $result .= chr($digit);
        }
    }
    return $result;
}

function string_to_hex(string $data): string {
    $chunks = [];
    for($i = 0; $i < strlen($data); $i += 4) {
        $chunk = "";
        for($s = 3; $s >= 0; $s -= 1) {
            $chunk .= value_to_digit(ord($data[$i + $s]) / 16);
            $chunk .= value_to_digit(ord($data[$i + $s]) % 16);
        }
        $chunks []= $chunk;
    }
    return join(" ", $chunks);
}

$PATH_TO_DATA = "../data/test-objects-bytes.json";

$value = json_decode(file_get_contents($PATH_TO_DATA));
$factory = new VK\TL\tl_factory();

foreach($value->Tests as $test_name => $test_data) {
    $testing_type = $test_data->TestingType;
    foreach($test_data->Successes as $exact_test) {
        $input_data = hex_to_string($exact_test->Bytes);

        $stream = new VK\TL\tl_input_stream($input_data);
        $output = new VK\TL\tl_output_stream("");

        $object = $factory->tl_object_by_name($testing_type);
        if ($object == null) {
            throw new Exception(sprintf("no such type \"%s\"", $testing_type));
        }

        $success = $object->read($stream);
        if (!$success) {
            throw new Exception(sprintf("\nunsuccessfull read for type \"%s\" on data \"%s\"", $testing_type, $exact_test["Bytes"]));
        }

        $success = $object->write($output);
        if (!$success) {
            throw new Exception(sprintf("\nunsuccessfull write for type \"%s\" on data \"%s\"", $testing_type, $exact_test["Bytes"]));
        }

        if ($input_data != $output->get_data()) {
            throw new Exception(sprintf("\nread and write result are different on \"%s\":\n\texpected: %s\n\tactual:   %s\n", $testing_type, string_to_hex($input_data), string_to_hex($output->get_data())));
        }

        printf("Test with \"%s\" passed on \"%s\"\n", $testing_type, string_to_hex($input_data));
    }
}

printf("\nAll tests are passed!\n");