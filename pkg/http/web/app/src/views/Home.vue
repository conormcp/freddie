<template>
  <layout>
    <v-carousel>
      <v-carousel-item
        v-for="(item, i) in items"
        :key="i"
        :src="item.src"
      ></v-carousel-item>
    </v-carousel>
    <v-container fluid>
      <v-layout row wrap justify-space-around>
        <v-flex xs12 sm11 md5> <meetings-box /> </v-flex>
        <v-flex xs12 sm11 md5> <papers-box /> </v-flex>
      </v-layout>
    </v-container>
  </layout>
</template>

<script>
import Layout from "@layouts/main";
import MeetingsBox from "@components/MeetingsBox";
import PapersBox from "@components/PapersBox";

export default {
  components: {
    Layout,
    MeetingsBox,
    PapersBox
  },
  data() {
    return {
      items: [
        {
          src:
            "https://www.nasa.gov/images/content/223974main_wildgalaxies1_20080424_HI.jpg"
        },
        {
          src:
            "http://keckobservatory.org/wp-content/uploads/2018/11/STSCI-H-p1828a-f-2422x2117.jpg"
        },
        {
          src:
            "https://www.nasa.gov/sites/default/files/images/597502main_hs-2008-16-aa-xlarge_web_full.jpg"
        },
        {
          src: "https://apod.nasa.gov/apod/image/1403/antennae_coverta_1824.jpg"
        }
      ]
    };
  }
};
</script>
