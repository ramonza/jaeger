// Copyright (c) 2018 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package static

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"

	"github.com/jaegertracing/jaeger/pkg/config"
	ss "github.com/jaegertracing/jaeger/pkg/sampling/strategystore"
	"github.com/jaegertracing/jaeger/plugin"
)

var _ ss.Factory = new(Factory)
var _ plugin.Configurable = new(Factory)

func TestFactory(t *testing.T) {
	f := NewFactory()
	v, command := config.Viperize(f.AddFlags)
	command.ParseFlags([]string{"--sampling.strategies=fixtures/strategies.json"})
	f.InitFromViper(v)

	assert.NoError(t, f.Initialize(metrics.NullFactory, zap.NewNop()))
	_, err := f.CreateStrategyStore()
	assert.NoError(t, err)
}