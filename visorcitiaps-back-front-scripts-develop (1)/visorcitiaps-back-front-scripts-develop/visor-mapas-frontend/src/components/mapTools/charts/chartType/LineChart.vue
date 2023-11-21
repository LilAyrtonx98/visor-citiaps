<script>
import { Line } from 'vue-chartjs'

export default {
  extends: Line,
  data () {
    return {
      gradient: null
    }
  },
  methods : {
    // Overwriting base render method with actual data.
    showData: function(chartData) {
      /*
      this.gradient = this.$refs.canvas.getContext('2d').createLinearGradient(0, 0, 0, 500)
      
      this.gradient.addColorStop(0, 'rgba(0, 231, 255, 0.5)')
      this.gradient.addColorStop(0.2, 'rgba(0, 231, 255, 0.2)');
      this.gradient.addColorStop(1, 'rgba(0, 231, 255, 0)');
      
      chartData.datasets[0].backgroundColor = this.gradient;
      */

      this.renderChart(
      chartData,
      chartData.options)
    }
  }
}
</script>
