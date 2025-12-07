// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {useHistory} from 'react-router-dom';
import {useIntl} from 'react-intl';

import PingLogoImage from 'images/ping-logo.png';

import './landing.scss';

const Landing = () => {
    const history = useHistory();
    const {formatMessage} = useIntl();

    const handleLoginClick = () => {
        history.push('/login');
    };

    return (
        <div className='landing-page'>
            <div className='landing-page__hero'>
                <div className='landing-page__content'>
                    <div className='landing-page__logo'>
                        <img
                            src={PingLogoImage}
                            alt='Ping'
                            className='landing-page__logo-image'
                        />
                    </div>
                    <h1 className='landing-page__title'>
                        {formatMessage({
                            id: 'landing.title',
                            defaultMessage: 'Secure Communication Platform',
                        })}
                    </h1>
                    <p className='landing-page__subtitle'>
                        {formatMessage({
                            id: 'landing.subtitle',
                            defaultMessage: 'Connect, collaborate, and communicate securely with your team.',
                        })}
                    </p>
                    <button
                        className='landing-page__login-button'
                        onClick={handleLoginClick}
                    >
                        {formatMessage({
                            id: 'landing.login',
                            defaultMessage: 'Login',
                        })}
                    </button>
                </div>
            </div>
        </div>
    );
};

export default Landing;

