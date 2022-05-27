import React from 'react';
import styles from './index.css';
import { DatePicker } from 'antd';


const App: React.FC = () => {
  return (
    <div className={styles.normal}>
      <div className={styles.welcome} />
      <ul className={styles.list}>
        <li>Hello, world! Welcome to Reikai, nothing to see here yet.</li>
      </ul>
    </div>
  );
}

export default App;
