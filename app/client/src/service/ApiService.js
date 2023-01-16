import axios from "axios";

const url = process.env.VITE_VUE_APP_API_URL;

axios.defaults.url = url;


