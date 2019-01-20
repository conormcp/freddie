<template>
  <v-toolbar app clipped fixed>
    <v-toolbar-title class="headline text-uppercase">
      <span>FREDDIE</span>
    </v-toolbar-title>
    <v-spacer></v-spacer>
    <v-toolbar-items>
      <v-btn :to="{ name: 'home' }" flat>Home</v-btn>
      <v-btn :to="{ name: 'calendar' }" flat>Calendar</v-btn>
      <v-btn :to="{ name: 'papers' }" flat>Papers</v-btn>
      <v-btn :to="{ name: 'about' }" flat>About</v-btn>
    </v-toolbar-items>
  </v-toolbar>
</template>
