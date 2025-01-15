<?php

/**
 * AUTOGENERATED, DO NOT EDIT! If you want to modify it, check tl schema.
 *
 * This autogenerated code represents tl class for typed RPC API.
 */

namespace VK\TL\cases_bytes\Types;

use VK\TL;

/**
 * @kphp-tl-class
 */
class cases_bytes_testEnumContainer implements TL\Readable, TL\Writeable {

  /** @var TL\cases\Types\cases_TestEnum */
  public $value = null;

  /**
   * @param TL\cases\Types\cases_TestEnum $value
   */
  public function __construct($value = null) {
    $this->value = $value;
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read_boxed($stream) {
    [$magic, $success] = $stream->read_uint32();
    if (!$success || $magic != 0x32b92037) {
      return false;
    }
    return $this->read($stream);
  }

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read($stream) {
    [$tag, $success] = $stream->read_uint32();
    if (!$success) {
      return false;
    }
    switch ($tag) {
      case 0x6c6c55ac:
        $variant = new TL\cases\Types\cases_testEnum1();
        $success = $variant->read($stream);
        if (!$success) {
          return false;
        }
        $this->value = $variant;
        break;
      case 0x86ea88ce:
        $variant = new TL\cases\Types\cases_testEnum2();
        $success = $variant->read($stream);
        if (!$success) {
          return false;
        }
        $this->value = $variant;
        break;
      case 0x69b83e2f:
        $variant = new TL\cases\Types\cases_testEnum3();
        $success = $variant->read($stream);
        if (!$success) {
          return false;
        }
        $this->value = $variant;
        break;
      default:
        return false;
    }
    return true;
  }

  /**
   * @param TL\tl_output_stream $stream
   * @return bool 
   */
  public function write_boxed($stream) {
    $success = $stream->write_uint32(0x32b92037);
    if (!$success) {
      return false;
    }
    return $this->write($stream);
  }

  /**
   * @param TL\tl_output_stream $stream
   * @return bool 
   */
  public function write($stream) {
    if ($this->value == null) {
      $this->value = new TL\cases\Types\cases_testEnum1();
    }
    $success = $this->value->write_boxed($stream);
    if (!$success) {
      return false;
    }
    return true;
  }

}
