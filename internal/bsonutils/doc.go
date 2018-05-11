// Copyright 2009 Dênnis Dantas de Sousa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package bsonutils reimplement BSON Marshal and Unmarshal from mgo package.

I've created this package to implement my needed version of Marshal and
Unmarshal functions: one that allows to set OmitEmpty tag as default.
Since it's code exclusive for use on this package, it was put as a
internal package.
*/
package bsonutils
