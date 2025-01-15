<?php

namespace VK\TL;

use VK\TL;

interface Readable {
  /**
   * @param TL\tl_input_stream $stream
   * @return bool
   */
  public function read($stream);

  /**
   * @param TL\tl_input_stream $stream
   * @return bool
   */
  public function read_boxed($stream);
}

interface Writeable {
  /**
   * @param TL\tl_output_stream $stream
   * @return bool
   */
  public function write($stream);

  /**
   * @param TL\tl_output_stream $stream
   * @return bool
   */
  public function write_boxed($stream);
}

interface TL_Object extends Readable, Writeable {}
