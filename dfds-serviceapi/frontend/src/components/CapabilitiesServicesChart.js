import {Pie} from 'vue-chartjs';


export default {
  extends: Pie,
  props: {
    chartdata: {
      type: Object,
      default: {}
    },
    options: {
      type: Object,
      default: null
    }
  },

  watch: {
    chartdata: function (newChart) {
      this.renderChart(newChart, this.options);
    }
  },

  mounted () {
    console.log("mounted");
    this.renderChart(this.chartdata, this.options)
  },

  created() {
    console.log("created");
  },

  beforeUpdate() {
    console.log("Updated");
    this.renderChart(this.chartdata, this.options)
  }
}
