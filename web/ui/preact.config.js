// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { EnvironmentPlugin } from 'webpack';

export default {
  webpack(config, env, helpers, options) {
    // Provide key environment variables we need to evaluate in our code.
    // The default values are used if the values are not already set in the
    // environment.
    config.plugins.push(
      new EnvironmentPlugin({
        NODE_ENV: 'development',
        DEBUG: false,
        FAKE_API: false,
      })
    );
  },
}
