<?php

namespace Dimitriin\Metrics\PrometheusLib;

interface LabelsResolverInterface
{
    /**
     * @param array $labelsMap
     * @return string[]
     */
    public function getLabelValues(array $labelsMap): array;
}