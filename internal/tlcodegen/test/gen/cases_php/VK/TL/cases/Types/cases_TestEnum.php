<?php

/**
 * AUTOGENERATED, DO NOT EDIT! If you want to modify it, check tl schema.
 *
 * This autogenerated code represents tl class for typed RPC API.
 */

namespace VK\TL\cases\Types;

use VK\TL;

/**
 * @kphp-tl-class
 */
interface cases_TestEnum {

  /** Allows kphp implicitly load all available constructors */
  const CONSTRUCTORS = [
    TL\cases\Types\cases_testEnum1::class,
    TL\cases\Types\cases_testEnum2::class,
    TL\cases\Types\cases_testEnum3::class
  ];

  /**
   * @param TL\tl_input_stream $stream
   * @return bool 
   */
  public function read_boxed($stream);

  /**
   * @param TL\tl_output_stream $stream
   * @return bool 
   */
  public function write_boxed($stream);
}
