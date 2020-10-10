<?php
declare(strict_types=1);

namespace Dimitriin\Probes;

interface ReadyBitSenderInterface
{
    public function send(): void;
}