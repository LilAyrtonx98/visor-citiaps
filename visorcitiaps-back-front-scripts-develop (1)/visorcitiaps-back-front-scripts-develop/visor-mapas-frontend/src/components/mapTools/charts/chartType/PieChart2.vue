<template>
  <div class="chart-container" v-if="chartData.isData">
    <ChartStat
    v-bind:chartContainerAlt="showAltView"
    v-bind:title="chartData.extra.title"
    v-bind:subtitle="chartData.extra.subtitle"
    v-bind:stats="chartData.extra.stats"
    >
    </ChartStat>
    <apexchart width="100%" type="pie" :options="chartData.chartOptions" :series="chartData.series"></apexchart>
  </div>
  <div v-else>
    <ChartStatNoData
    v-bind:chartContainerAlt="showAltView"
    v-bind:type="0"/>
  </div>
</template>

<script>
import ChartStat from '@/components/mapTools/charts/ChartStat.vue';
import ChartStatNoData from '@/components/mapTools/charts/ChartStatNoData.vue';

export default {
    props: ['chartData', 'showAltView'],
    components: {
      ChartStat,
      ChartStatNoData
    },
    data: function() {
      return {

      }
    },
    methods: {

    }
};
</script>
