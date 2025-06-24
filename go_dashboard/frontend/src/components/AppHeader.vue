<template>
  <header class="header">
    <div class="logo">Dashboard</div>

    <nav class="nav" :class="{ open: isMenuOpen }">
      <a href="#">Home</a>
      <a href="#">Reports</a>
      <a href="#">Settings</a>
    </nav>

    <button class="menu-toggle" @click="toggleMenu">
      ☰
    </button>
  </header>
</template>

<script>
export default {
  data() {
    return {
      isMenuOpen: false
    };
  },
  methods: {
    toggleMenu() {
      this.isMenuOpen = !this.isMenuOpen;
    }
  }
};
</script>

<style scoped>
.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #2c3e50;
  padding: 1rem;
  color: white;
  position: relative;
  flex-wrap: wrap;
}

.logo {
  font-size: 1.5rem;
  font-weight: bold;
}

.nav {
  display: flex;
  gap: 1.5rem;
}

.nav a {
  color: white;
  text-decoration: none;
  font-size: 1rem;
}

.menu-toggle {
  font-size: 1.5rem;
  background: none;
  color: white;
  border: none;
  display: none;
  cursor: pointer;
}

/* Responsive styles */
@media (max-width: 768px) {
  .menu-toggle {
    display: block;
  }

  .nav {
    display: none;
    flex-direction: column;
    width: 100%;
    margin-top: 1rem;
    background-color: #34495e;
    padding: 1rem;
    border-radius: 4px;
  }

  .nav.open {
    display: flex;
  }

  .nav a {
    margin: 0.5rem 0;
  }
}
</style>
