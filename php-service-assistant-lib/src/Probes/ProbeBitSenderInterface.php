<?php
declare(strict_types=1);

namespace Dimitriin\Probes;

interface ProbeBitSenderInterface
{
    public function send(?int $ttl = null): void;
}