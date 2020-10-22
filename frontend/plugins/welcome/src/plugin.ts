import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import login from './components/login';
import table from './components/Table';
import home from './components/home';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/home', home);
    router.registerRoute('/welcome', WelcomePage);
    router.registerRoute('/', login);
    router.registerRoute('/table', table);
  },
});
