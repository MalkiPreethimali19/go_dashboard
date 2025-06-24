<template>
  <div class="card">
    <h2>Top 30 Regions by Revenue</h2>
    <div class="chart-wrapper">
      <Bar :data="chartData" :options="chartOptions" v-if="loaded" />
    </div>
  </div>
</template>

<script>
import { Bar } from 'vue-chartjs'
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  CategoryScale,
  LinearScale,
  BarElement
} from 'chart.js'
import api from '../api'

ChartJS.register(Title, Tooltip, Legend, CategoryScale, LinearScale, BarElement)

export default {
  components: { Bar },
  data() {
    return {
      chartData: {
        labels: [],
        datasets: []
      },
      chartOptions: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: { display: false },
          tooltip: {
            callbacks: {
              label: (ctx) => `Revenue: $${ctx.raw.toFixed(2)}`
            }
          }
        },
        scales: {
          y: {
            beginAtZero: true,
            title: { display: true, text: 'Total Revenue ($)' }
          },
          x: {
            ticks: { autoSkip: false }
          }
        }
      },
      loaded: false
    }
  },
  mounted() {
    api.get('/top-regions').then(res => {
      const labels = res.data.map(r => `${r.region_name}`)
      const values = res.data.map(r => r.total_revenue)

      this.chartData = {
        labels,
        datasets: [{
          label: 'Revenue',
          data: values,
          backgroundColor: this.generateColors(values.length)
        }]
      }

      this.loaded = true
    }).catch(err => {
      console.error('Error loading region revenue:', err)
    })
  },
  methods: {
    generateColors(count) {
      return Array.from({ length: count }, (_, i) => `hsl(${(i * 360) / count}, 70%, 60%)`)
    }
  }
}
</script>

<style scoped>
.card {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  margin-bottom: 2rem;
  box-shadow: 0 2px 10px rgba(0,0,0,0.05);
}

.chart-wrapper {
  width: 100%;
  height: 600px;
  overflow-x: auto;
}
</style>
