import React from 'react';
import AdSense from 'react-adsense';

export default class AdsenseWidget extends React.Component {
	componentDidMount() {
		const installGoogleAds = () => {
			const elem = document.createElement("script");
			elem.src = "//pagead2.googlesyndication.com/pagead/js/adsbygoogle.js";
			elem.async = true;
			elem.defer = true;
			document.body.insertBefore(elem, document.body.firstChild);
		};
		installGoogleAds();
	}


render () {
	return (
		<AdSense.Google
			client='ca-pub-6526798120378756'
			slot='1686195266'
			style={{ display: 'block' }}
			format='auto'
			responsive='true'
		/>
	);
	}
}
