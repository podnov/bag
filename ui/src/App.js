import React from 'react';
import logo from './logo.svg';
import './App.css';
import WalletInput from './components/wallet/WalletInput';
import WalletStatistics from './components/wallet/WalletStatistics';
import AdsenseWidget from './components/adsense/AdsenseWidget';

function App() {

	return (
		<div className="App">
			<header>
				<h1>CryptoBag</h1>
			</header>
			<WalletInput />
			<div className="content">
				<WalletStatistics />
			</div>
			<footer>
				<div>
					Powered by:&nbsp;
					<a href="https://bscscan.com" rel="noreferrer" target="_blank">BscScan.com</a> APIs,&nbsp;
					<a href="https://github.com/pancakeswap/pancake-info-api" rel="noreferrer" target="_blank">Pancakeswap APIs</a>,&nbsp;
					<a href="https://github.com" rel="noreferrer" target="_blank">GitHub</a>
				</div>
				<AdsenseWidget />
			</footer>
		</div>
	);
}

export default App;