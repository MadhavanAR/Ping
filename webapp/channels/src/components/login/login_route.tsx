// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import {Route} from 'react-router-dom';

import AnnouncementBar from 'components/announcement_bar';

import Login from './login';

import './login_route.scss';

const LoginRoute = ({path}: {path: string}) => {
    return (
        <Route
            path={path}
            render={() => (
                <>
                    <React.Suspense fallback={null}>
                        <AnnouncementBar/>
                    </React.Suspense>
                    <div className='login-route'>
                        <Login/>
                    </div>
                </>
            )}
        />
    );
};

export default LoginRoute;

