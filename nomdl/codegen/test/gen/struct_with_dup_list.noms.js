// This file was generated by nomdl/codegen.
// @flow
/* eslint-disable */

import {
  Field as _Field,
  Kind as _Kind,
  Package as _Package,
  createStructClass as _createStructClass,
  makeCompoundType as _makeCompoundType,
  makeListType as _makeListType,
  makeStructType as _makeStructType,
  makeType as _makeType,
  newList as _newList,
  registerPackage as _registerPackage,
  uint8Type as _uint8Type,
} from '@attic/noms';
import type {
  NomsList as _NomsList,
  Struct as _Struct,
  uint8 as _uint8,
} from '@attic/noms';

const _pkg = new _Package([
  _makeStructType('StructWithDupList',
      [
        new _Field('l', _makeCompoundType(_Kind.List, _uint8Type), false),
      ],
      [

      ]
    ),
], [
]);
_registerPackage(_pkg);
export const typeForStructWithDupList = _makeType(_pkg.ref, 0);
const StructWithDupList$typeDef = _pkg.types[0];


type StructWithDupList$Data = {
  l: _NomsList<_uint8>;
};

interface StructWithDupList$Interface extends _Struct {
  constructor(data: StructWithDupList$Data): void;
  l: _NomsList<_uint8>;  // readonly
  setL(value: _NomsList<_uint8>): StructWithDupList$Interface;
}

export const StructWithDupList: Class<StructWithDupList$Interface> = _createStructClass(typeForStructWithDupList, StructWithDupList$typeDef);

export function newListOfUint8(values: Array<_uint8>): Promise<_NomsList<_uint8>> {
  return _newList(values, _makeListType(_uint8Type));
}
