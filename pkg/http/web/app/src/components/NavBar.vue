<template>
  <v-toolbar app clipped fixed>
    <v-toolbar-title class="headline text-uppercase">
      <span>FREDDIE</span>
    </v-toolbar-title>
    <v-spacer></v-spacer>
    <v-toolbar-items v-if="$vuetify.breakpoint.mdAndUp">
      <v-btn :to="{ name: 'home' }" flat>Home</v-btn>
      <v-btn :to="{ name: 'calendar' }" flat>Calendar</v-btn>
      <v-btn :to="{ name: 'papers' }" flat>Papers</v-btn>
    </v-toolbar-items>
    <v-menu v-else>
      <v-btn slot="activator" icon><v-icon>menu</v-icon></v-btn>
      <v-list>
        <v-list-tile :to="{ name: 'home' }">
          <v-list-tile-title>Home</v-list-tile-title>
        </v-list-tile>
        <v-list-tile :to="{ name: 'calendar' }">
          <v-list-tile-title>Calendar</v-list-tile-title>
        </v-list-tile>
        <v-list-tile :to="{ name: 'papers' }">
          <v-list-tile-title>Papers</v-list-tile-title>
        </v-list-tile>
      </v-list>
    </v-menu>
  </v-toolbar>
</template>
