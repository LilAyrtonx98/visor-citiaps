<script>
import { Pie } from 'vue-chartjs'

export default {
  extends: Pie,
  methods : {
    // Overwriting base render method with actual data.
    showData: function(chartData) {
      this.renderChart(
      chartData,
      chartData.options)
    }
  }
}
</script>
