import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import 'primeicons/primeicons.css'
import PrimeVue from 'primevue/config';
import Aura from '@primeuix/themes/aura';

const app = createApp(App)
app.use(PrimeVue, {
    theme: {
        preset: Aura,
        options: {
            darkModeSelector: ".p-dark",
        },
    }
})
app.mount('#app')
