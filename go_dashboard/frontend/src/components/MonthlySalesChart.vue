<template>
  <div class="card">
    <h2>Monthly Sales Volume</h2>
    <div class="chart-wrapper">
      <Line :data="chartData" :options="chartOptions" v-if="loaded" />
    </div>
  </div>
</template>

<script>
import { Line } from 'vue-chartjs'
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  LineElement,
  PointElement,
  CategoryScale,
  LinearScale,
} from 'chart.js'
import api from '../api'

ChartJS.register(Title, Tooltip, Legend, LineElement, PointElement, CategoryScale, LinearScale)

export default {
  components: { Line },
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
          legend: {
            display: false
          },
          tooltip: {
            callbacks: {
              label: (ctx) => `Transactions: ${ctx.raw}`
            }
          }
        },
        scales: {
          y: {
            beginAtZero: true,
            title: {
              display: true,
              text: 'Transaction Count'
            }
          },
          x: {
            title: {
              display: true,
              text: 'Month'
            }
          }
        }
      },
      loaded: false
    }
  },
  mounted() {
    api.get('/monthly-sales').then((res) => {
      const labels = res.data.map(d => d.month)
      const values = res.data.map(d => d.transaction_count)

      this.chartData = {
        labels,
        datasets: [{
          label: 'Transactions',
          data: values,
          fill: false,
          borderColor: '#42A5F5',
          backgroundColor: '#42A5F5',
          tension: 0.3
        }]
      }
      this.loaded = true
    }).catch(err => {
      console.error("Error loading monthly sales:", err)
    })
  }
}
</script>

<style scoped>
.card {
  background: white;
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 10px rgba(0,0,0,0.05);
}

.chart-wrapper {
  width: 100%;
  height: 400px;
  overflow-x: auto;
}
</style>
