import "package:bloc/bloc.dart";
import "package:booking_client/api/host_api.dart";
import "package:booking_client/models/host.dart";
import "package:booking_client/utils/result.dart";
import "package:flutter/foundation.dart";

@immutable
sealed class HostsState {
  const HostsState({required this.hosts});

  final List<Host>? hosts;

  HostsState copyWith({List<Host>? hosts});
}

class HostsInitial extends HostsState {
  const HostsInitial({required super.hosts});

  @override
  HostsState copyWith({List<Host>? hosts}) =>
      HostsInitial(hosts: hosts ?? this.hosts);
}

class HostsLoading extends HostsState {
  const HostsLoading({required super.hosts});

  @override
  HostsState copyWith({List<Host>? hosts}) =>
      HostsLoading(hosts: hosts ?? this.hosts);
}

class HostsSuccess extends HostsState {
  const HostsSuccess({required super.hosts});

  @override
  HostsState copyWith({List<Host>? hosts}) =>
      HostsSuccess(hosts: hosts ?? this.hosts);
}

class HostsFailure extends HostsState {
  const HostsFailure({required super.hosts, required this.error});

  final ApiError error;

  @override
  HostsState copyWith({List<Host>? hosts}) =>
      HostsFailure(hosts: hosts ?? this.hosts, error: error);
}

class HostsCubit extends Cubit<HostsState> {
  HostsCubit(this._hostApi) : super(const HostsInitial(hosts: []));

  final HostApi _hostApi;

  Future<void> getHosts() async {
    emit(HostsLoading(hosts: state.hosts));

    final result = await _hostApi.getHosts();
    if (result.isFailure) {
      emit(HostsFailure(hosts: state.hosts, error: result.error));
    } else {
      emit(HostsSuccess(hosts: result.value));
    }
  }
}
